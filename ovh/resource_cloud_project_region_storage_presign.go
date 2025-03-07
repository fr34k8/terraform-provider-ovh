package ovh

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/v2/ovh/helpers"
)

func resourceCloudProjectRegionStoragePresign() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudProjectRegionStoragePresignCreate,
		Read:   schema.Noop,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				DefaultFunc: schema.EnvDefaultFunc("OVH_CLOUD_PROJECT_SERVICE", nil),
				Description: "Service name of the resource representing the ID of the cloud project",
			},
			"region_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Region name",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The S3 storage container's name",
			},
			"expire": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "How long (in seconds) the URL will be valid",
			},
			"method": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: func(v interface{}, k string) (ws []string, errors []error) {
					err := helpers.ValidateStringEnum(v.(string), []string{"GET", "PUT"})
					if err != nil {
						errors = append(errors, err)
					}
					return
				},
			},
			"object": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the object to download or upload",
			},
			"version_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Version ID of the object to download or delete",
			},

			// Computed
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Presigned URL",
			},
			"signed_headers": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "Signed headers",
			},
		},
	}
}

func resourceCloudProjectRegionStoragePresignCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	regionName := strings.ToUpper(d.Get("region_name").(string))
	name := d.Get("name").(string)

	resp := &PresignedURL{}
	opts := (&PresignedURLInput{}).FromResource(d)

	endpoint := fmt.Sprintf("/cloud/project/%s/region/%s/storage/%s/presign", url.PathEscape(serviceName), url.PathEscape(regionName), url.PathEscape(name))
	if err := config.OVHClient.Post(endpoint, opts, resp); err != nil {
		return fmt.Errorf("error calling post %s:\n\t %q", endpoint, err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	if err := d.Set("url", resp.URL); err != nil {
		return err
	}

	if err := d.Set("signed_headers", resp.SignedHeaders); err != nil {
		return err
	}

	return nil
}
