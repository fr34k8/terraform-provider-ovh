package ovh

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/ovh/helpers"
)

type DedicatedServer struct {
	AvailabilityZone   string `json:"availabilityZone"`
	Name               string `json:"name"`
	BootId             int    `json:"bootId"`
	BootScript         string `json:"bootScript"`
	CommercialRange    string `json:"commercialRange"`
	Datacenter         string `json:"datacenter"`
	EfiBootloaderPath  string `json:"efiBootloaderPath"`
	Ip                 string `json:"ip"`
	LinkSpeed          int    `json:"linkSpeed"`
	Monitoring         bool   `json:"monitoring"`
	NewUpgradeSystem   bool   `json:"newUpgradeSystem"`
	NoIntervention     bool   `json:"noIntervention"`
	Os                 string `json:"os"`
	PowerState         string `json:"powerState"`
	ProfessionalUse    bool   `json:"professionalUse"`
	Rack               string `json:"rack"`
	Region             string `json:"region"`
	RescueMail         string `json:"rescueMail"`
	RescueSshKey       string `json:"rescueSshKey"`
	Reverse            string `json:"reverse"`
	RootDevice         string `json:"rootDevice"`
	ServerId           int    `json:"serverId"`
	State              string `json:"state"`
	SupportLevel       string `json:"supportLevel"`
	IamResourceDetails `json:"iam"`
}

func (ds DedicatedServer) String() string {
	return fmt.Sprintf(
		"name: %v, ip: %v, dc: %v, state: %v",
		ds.Name,
		ds.Ip,
		ds.Datacenter,
		ds.State,
	)
}

type DedicatedServerUpdateOpts struct {
	BootId            *int64  `json:"bootId,omitempty"`
	BootScript        *string `json:"bootScript,omitempty"`
	EfiBootloaderPath *string `json:"efiBootloaderPath,omitempty"`
	Monitoring        *bool   `json:"monitoring,omitempty"`
	State             *string `json:"state,omitempty"`
}

func (opts *DedicatedServerUpdateOpts) FromResource(d *schema.ResourceData) *DedicatedServerUpdateOpts {
	opts.BootId = helpers.GetNilInt64PointerFromData(d, "boot_id")
	opts.BootScript = helpers.GetNilStringPointerFromData(d, "boot_script")
	opts.EfiBootloaderPath = helpers.GetNilStringPointerFromData(d, "efi_bootloader_path")
	opts.Monitoring = helpers.GetNilBoolPointerFromData(d, "monitoring")
	opts.State = helpers.GetNilStringPointerFromData(d, "state")
	return opts
}

type DedicatedServerVNI struct {
	Enabled    bool     `json:"enabled"`
	Mode       string   `json:"mode"`
	Name       string   `json:"name"`
	NICs       []string `json:"networkInterfaceController"`
	ServerName string   `json:"serverName"`
	Uuid       string   `json:"uuid"`
	Vrack      *string  `json:"vrack,omitempty"`
}

func (vni DedicatedServerVNI) String() string {
	return fmt.Sprintf(
		"name: %v, uuid: %v, mode: %v, enabled: %v",
		vni.Name,
		vni.Uuid,
		vni.Mode,
		vni.Enabled,
	)
}

func (v DedicatedServerVNI) ToMap() map[string]interface{} {
	obj := make(map[string]interface{})
	obj["enabled"] = v.Enabled
	obj["mode"] = v.Mode
	obj["name"] = v.Name
	obj["nics"] = v.NICs
	obj["server_name"] = v.ServerName
	obj["uuid"] = v.Uuid

	if v.Vrack != nil {
		obj["vrack"] = *v.Vrack
	}

	return obj
}

type DedicatedServerBoot struct {
	BootId      int    `json:"bootId"`
	BootType    string `json:"bootType"`
	Description string `json:"description"`
	Kernel      string `json:"kernel"`
}

type DedicatedServerTask struct {
	Id         int64     `json:"taskId"`
	Function   string    `json:"function"`
	Comment    string    `json:"comment"`
	Status     string    `json:"status"`
	LastUpdate time.Time `json:"lastUpdate"`
	DoneDate   time.Time `json:"doneDate"`
	StartDate  time.Time `json:"startDate"`
}

type DedicatedServerInstallTaskUserMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DedicatedServerInstallTaskCreateOpts struct {
	TemplateName        string                                   `json:"templateName"`
	PartitionSchemeName *string                                  `json:"partitionSchemeName,omitempty"`
	Details             *DedicatedServerInstallTaskDetails       `json:"details"`
	UserMetadata        []DedicatedServerInstallTaskUserMetadata `json:"userMetadata,omitempty"`
}

func (opts *DedicatedServerInstallTaskCreateOpts) FromResource(d *schema.ResourceData) *DedicatedServerInstallTaskCreateOpts {
	opts.TemplateName = d.Get("template_name").(string)
	opts.PartitionSchemeName = helpers.GetNilStringPointerFromData(d, "partition_scheme_name")

	details := d.Get("details").([]interface{})
	if len(details) == 1 {
		opts.Details = (&DedicatedServerInstallTaskDetails{}).FromResource(d, "details.0")
	}

	userMetadata := d.Get("user_metadata").([]interface{})
	var userMetadatas []DedicatedServerInstallTaskUserMetadata

	for _, metadata := range userMetadata {
		m := metadata.(map[string]interface{})
		key := m["key"].(string)
		value := m["value"].(string)
		metadatum := DedicatedServerInstallTaskUserMetadata{Key: key, Value: value}
		userMetadatas = append(userMetadatas, metadatum)
	}
	opts.UserMetadata = userMetadatas

	return opts
}

type DedicatedServerInstallTaskDetails struct {
	CustomHostname  *string `json:"customHostname,omitempty"`
	DiskGroupId     *int64  `json:"diskGroupId,omitempty"`
	NoRaid          *bool   `json:"noRaid,omitempty"`
	SoftRaidDevices *int64  `json:"softRaidDevices,omitempty"`
}

func (opts *DedicatedServerInstallTaskDetails) FromResource(d *schema.ResourceData, parent string) *DedicatedServerInstallTaskDetails {
	opts.CustomHostname = helpers.GetNilStringPointerFromData(d, fmt.Sprintf("%s.custom_hostname", parent))
	opts.DiskGroupId = helpers.GetNilInt64PointerFromData(d, fmt.Sprintf("%s.disk_group_id", parent))
	opts.NoRaid = helpers.GetNilBoolPointerFromData(d, fmt.Sprintf("%s.no_raid", parent))
	opts.SoftRaidDevices = helpers.GetNilInt64PointerFromData(d, fmt.Sprintf("%s.soft_raid_devices", parent))

	return opts
}
