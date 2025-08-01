package ovh

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/ovh/go-ovh/ovh"
	"github.com/ovh/terraform-provider-ovh/v2/ovh/helpers"
	"github.com/ovh/terraform-provider-ovh/v2/ovh/ovhwrap"
)

type OvhDomainZoneRecord struct {
	Id        int64  `json:"id,omitempty"`
	Zone      string `json:"zone,omitempty"`
	Target    string `json:"target"`
	Ttl       int    `json:"ttl"`
	FieldType string `json:"fieldType,omitempty"`
	SubDomain string `json:"subDomain,omitempty"`
}

func (r *OvhDomainZoneRecord) String() string {
	return fmt.Sprintf(
		"record[id: %v, zone: %s, subdomain: %s, type: %s, target: %s]",
		r.Id,
		r.Zone,
		r.SubDomain,
		r.FieldType,
		r.Target,
	)
}

func resourceOvhDomainZoneRecordImportState(
	d *schema.ResourceData,
	meta interface{}) ([]*schema.ResourceData, error) {
	givenId := d.Id()
	splitId := strings.SplitN(givenId, ".", 2)
	if len(splitId) != 2 {
		return nil, fmt.Errorf("Import Id is not OVH_ID.zone formatted")
	}
	d.SetId(splitId[0])
	d.Set("zone", splitId[1])
	results := make([]*schema.ResourceData, 1)
	results[0] = d
	return results, nil
}

func resourceOvhDomainZoneRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceOvhDomainZoneRecordCreate,
		Read:   resourceOvhDomainZoneRecordRead,
		Update: resourceOvhDomainZoneRecordUpdate,
		Delete: resourceOvhDomainZoneRecordDelete,
		Importer: &schema.ResourceImporter{
			State: resourceOvhDomainZoneRecordImportState,
		},

		Schema: map[string]*schema.Schema{
			"zone": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(int)
					if v < 60 && v != 0 {
						errs = append(errs, fmt.Errorf("%q must be either equal to 0 or, greater than or equal to 60, got: %d", key, v))
					}
					return
				},
			},
			"fieldtype": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subdomain": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceOvhDomainZoneRecordCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	zone := d.Get("zone").(string)

	// Create the new record
	newRecord := &OvhDomainZoneRecord{
		FieldType: d.Get("fieldtype").(string),
		SubDomain: d.Get("subdomain").(string),
		Target:    d.Get("target").(string),
		Ttl:       d.Get("ttl").(int),
	}

	log.Printf("[DEBUG] OVH Record create configuration: %#v", newRecord)

	resultRecord := &OvhDomainZoneRecord{}

	err := config.OVHClient.Post(
		fmt.Sprintf("/domain/zone/%s/record", zone),
		newRecord,
		resultRecord,
	)

	if err != nil {
		return fmt.Errorf("Failed to create OVH Record: %s", err)
	}

	// this is an API response BUG known by OVH team
	// with no planned fix
	// Workaround is to filter records matching the attributes
	// and keep the last id if there are duplicates
	if resultRecord.Id == 0 {
		log.Printf("[WARN] Known OVH API Bug with Inconsistency API result (id = 0): %v", resultRecord)
		records := make([]int, 0)
		if err := config.OVHClient.CallAPI("GET", fmt.Sprintf("/domain/zone/%s/record", zone), newRecord, &records, true); err != nil {
			return fmt.Errorf("Error calling /domain/zone/%s. Zone may have been left with orphan records!:\n\t %q", zone, err)
		}

		if len(records) == 0 {
			return fmt.Errorf("API inconsistency: record creation on zone %s didn't fail but unable to retrieve it.", zone)
		}
		// reverse order to keep the last item if found
		sort.Sort(sort.Reverse(sort.IntSlice(records)))
		for _, rec := range records {
			record, err := ovhDomainZoneRecord(config.OVHClient, d, strconv.Itoa(rec), true)
			if err != nil {
				return fmt.Errorf("Error calling /domain/zone/%s. Zone may have been left with orphan records!:\n\t %q", zone, err)
			}

			if record == nil {
				log.Printf("[DEBUG] record %v has been deleted.", record)
				continue
			}

			log.Printf("[DEBUG] record found %v", record)
			if record.Target == newRecord.Target &&
				record.SubDomain == newRecord.SubDomain &&
				record.FieldType == newRecord.FieldType {
				resultRecord = record
				continue
			}
		}

	}

	d.SetId(strconv.FormatInt(resultRecord.Id, 10))

	if err := ovhDomainZoneRefresh(d, meta); err != nil {
		log.Printf("[WARN] OVH Domain zone refresh after record creation failed: %s", err)
	}

	return resourceOvhDomainZoneRecordRead(d, meta)
}

func resourceOvhDomainZoneRecordRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	record, err := ovhDomainZoneRecord(config.OVHClient, d, d.Id(), d.IsNewResource())
	if err != nil {
		return err
	}

	if record == nil {
		// ID already set to "", so just return
		log.Print("record has been deleted")
		return nil
	}

	d.Set("zone", record.Zone)
	d.Set("fieldtype", record.FieldType)
	d.Set("subdomain", record.SubDomain)
	d.Set("ttl", record.Ttl)
	d.Set("target", record.Target)

	return nil
}

func resourceOvhDomainZoneRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	record := OvhDomainZoneRecord{}

	if attr, ok := d.GetOk("subdomain"); ok {
		record.SubDomain = attr.(string)
	}
	if attr, ok := d.GetOk("target"); ok {
		record.Target = attr.(string)
	}
	if attr, ok := d.GetOk("ttl"); ok {
		record.Ttl, _ = attr.(int)
	}

	log.Printf("[DEBUG] OVH Record update configuration: %#v", record)

	err := config.OVHClient.Put(
		fmt.Sprintf("/domain/zone/%s/record/%s", d.Get("zone").(string), d.Id()),
		record,
		nil,
	)

	if err != nil {
		return fmt.Errorf("Failed to update OVH Record: %s", err)
	}

	if err := ovhDomainZoneRefresh(d, meta); err != nil {
		log.Printf("[WARN] OVH Domain zone refresh after record update failed: %s", err)
	}

	return resourceOvhDomainZoneRecordRead(d, meta)
}

func resourceOvhDomainZoneRecordDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	log.Printf("[INFO] Deleting OVH Record: %s.%s, %s", d.Get("zone").(string), d.Get("subdomain").(string), d.Id())

	endpoint := fmt.Sprintf("/domain/zone/%s/record/%s", d.Get("zone").(string), d.Id())
	if err := config.OVHClient.Delete(endpoint, nil); err != nil {
		return helpers.CheckDeleted(d, err, endpoint)
	}

	d.SetId("")

	if err := ovhDomainZoneRefresh(d, meta); err != nil {
		log.Printf("[WARN] OVH Domain zone refresh after record deletion failed: %s", err)
	}

	return nil
}

func ovhDomainZoneRefresh(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	log.Printf("[INFO] Refresh OVH Zone: %s", d.Get("zone").(string))

	err := config.OVHClient.Post(
		fmt.Sprintf("/domain/zone/%s/refresh", d.Get("zone").(string)),
		nil,
		nil,
	)

	if err != nil {
		return fmt.Errorf("Error refresh OVH Zone: %s", err)
	}

	return nil
}

func ovhDomainZoneRecord(client *ovhwrap.Client, d *schema.ResourceData, id string, retry bool) (*OvhDomainZoneRecord, error) {
	rec := &OvhDomainZoneRecord{}
	zone := d.Get("zone").(string)

	endpoint := fmt.Sprintf("/domain/zone/%s/record/%s", zone, id)

	err := resource.Retry(1*time.Minute, func() *resource.RetryError {
		err := client.Get(
			endpoint,
			rec,
		)
		if err != nil {
			if errOvh, ok := err.(*ovh.APIError); retry && ok && errOvh.Code == 404 {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})

	if err != nil {
		return nil, helpers.CheckDeleted(d, err, endpoint)
	}

	return rec, err
}
