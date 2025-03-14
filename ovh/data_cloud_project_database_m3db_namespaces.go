package ovh

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ovh/terraform-provider-ovh/v2/ovh/helpers/hashcode"
)

func dataSourceCloudProjectDatabaseM3dbNamespaces() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudProjectDatabaseM3dbNamespacesRead,
		Schema: map[string]*schema.Schema{
			"service_name": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OVH_CLOUD_PROJECT_SERVICE", nil),
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Description: "Cluster ID",
				Required:    true,
			},

			//Computed
			"namespace_ids": {
				Type:        schema.TypeList,
				Description: "List of namespaces ids",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceCloudProjectDatabaseM3dbNamespacesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)
	serviceName := d.Get("service_name").(string)
	clusterId := d.Get("cluster_id").(string)

	endpoint := fmt.Sprintf("/cloud/project/%s/database/m3db/%s/namespace",
		url.PathEscape(serviceName),
		url.PathEscape(clusterId),
	)

	res := make([]string, 0)

	log.Printf("[DEBUG] Will read namespaces from cluster %s from project %s", clusterId, serviceName)
	if err := config.OVHClient.GetWithContext(ctx, endpoint, &res); err != nil {
		return diag.Errorf("Error calling GET %s:\n\t %q", endpoint, err)
	}

	// sort.Strings sorts in place, returns nothing
	sort.Strings(res)

	d.SetId(hashcode.Strings(res))
	d.Set("namespace_ids", res)

	return nil
}
