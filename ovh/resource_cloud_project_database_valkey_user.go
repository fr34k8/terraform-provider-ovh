package ovh

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/v2/ovh/helpers"
)

func resourceCloudProjectDatabaseValkeyUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudProjectDatabaseValkeyUserCreate,
		ReadContext:   resourceCloudProjectDatabaseValkeyUserRead,
		DeleteContext: resourceCloudProjectDatabaseValkeyUserDelete,
		UpdateContext: resourceCloudProjectDatabaseValkeyUserUpdate,

		Importer: &schema.ResourceImporter{
			State: resourceCloudProjectDatabaseValkeyUserImportState,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OVH_CLOUD_PROJECT_SERVICE", nil),
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Description: "Id of the database cluster",
				ForceNew:    true,
				Required:    true,
			},
			"categories": {
				Type:        schema.TypeSet,
				Description: "Categories of the user",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"commands": {
				Type:        schema.TypeSet,
				Description: "Commands of the user",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"keys": {
				Type:        schema.TypeSet,
				Description: "Keys of the user",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the user",
				ForceNew:    true,
				Required:    true,
			},
			"password_reset": {
				Type:        schema.TypeString,
				Description: "Arbitrary string to change to trigger a password update",
				Optional:    true,
			},

			//Optional/Computed
			"channels": {
				Type:        schema.TypeSet,
				Description: "Channels of the user",
				Optional:    true,
				// If no channels list, channels = ["*"] is computed at creation
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},

			//Computed
			"created_at": {
				Type:        schema.TypeString,
				Description: "Date of the creation of the user",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Password of the user",
				Sensitive:   true,
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "Current status of the user",
				Computed:    true,
			},
		},

		CustomizeDiff: customdiff.ComputedIf(
			"password",
			func(ctx context.Context, d *schema.ResourceDiff, meta interface{}) bool {
				return d.HasChange("password_reset")
			},
		),
	}
}

func resourceCloudProjectDatabaseValkeyUserImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return importCloudProjectDatabaseUser(d, meta)
}

func resourceCloudProjectDatabaseValkeyUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	f := func() interface{} {
		return (&CloudProjectDatabaseRedisUserCreateOpts{}).FromResource(d)
	}
	return postCloudProjectDatabaseUser(ctx, d, meta, "valkey", dataSourceCloudProjectDatabaseValkeyUserRead, resourceCloudProjectDatabaseValkeyUserRead, resourceCloudProjectDatabaseValkeyUserUpdate, f)
}

func resourceCloudProjectDatabaseValkeyUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	clusterId := d.Get("cluster_id").(string)
	id := d.Id()

	endpoint := fmt.Sprintf("/cloud/project/%s/database/valkey/%s/user/%s",
		url.PathEscape(serviceName),
		url.PathEscape(clusterId),
		url.PathEscape(id),
	)
	res := &CloudProjectDatabaseRedisUserResponse{}

	log.Printf("[DEBUG] Will read user %s from cluster %s from project %s", id, clusterId, serviceName)
	if err := config.OVHClient.GetWithContext(ctx, endpoint, res); err != nil {
		return diag.FromErr(helpers.CheckDeleted(d, err, endpoint))
	}

	for k, v := range res.ToMap() {
		if k != "id" {
			d.Set(k, v)
		} else {
			d.SetId(fmt.Sprint(v))
		}
	}

	return nil
}

func resourceCloudProjectDatabaseValkeyUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	f := func() interface{} {
		return (&CloudProjectDatabaseRedisUserUpdateOpts{}).FromResource(d)
	}
	return updateCloudProjectDatabaseUser(ctx, d, meta, "valkey", resourceCloudProjectDatabaseValkeyUserRead, f)
}

func resourceCloudProjectDatabaseValkeyUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return deleteCloudProjectDatabaseUser(ctx, d, meta, "valkey")
}
