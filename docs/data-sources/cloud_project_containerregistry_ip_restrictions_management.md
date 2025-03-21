---
subcategory : "Managed Private Registry (MPR)"
---

# ovh_cloud_project_containerregistry_ip_restrictions_management (Data Source)

Use this data source to get the list of Management IP Restrictions of a container registry associated with a public cloud project.

## Example Usage

```terraform
data data "ovh_cloud_project_containerregistry_ip_restrictions_management" "mgt_iprestrictions_data" {
  service_name = "XXXXXX"
  registry_id  = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"
}

output "my_ip_restrictions" {
  value = data.ovh_cloud_project_containerregistry_ip_restrictions_management.mgt_iprestrictions_data.ip_restrictions
}
```

## Argument Reference

The following arguments are supported:

* `service_name` - (Optional) The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
* `registry_id` - The id of the Managed Private Registry.

## Attributes Reference

The following attributes are exported:

* `service_name` - The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
* `registry_id` - The ID of the Managed Private Registry.
* `ip_restrictions` - IP restrictions applied on Harbor UI and API.
  * `description` - The Description of Whitelisted IpBlock.
  * `ip_block` - Whitelisted IpBlock (CIDR format).
