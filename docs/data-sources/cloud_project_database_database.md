---
subcategory : "Managed Databases"
---

# ovh_cloud_project_database_database (Data Source)

Use this data source to get information about a database of a database cluster associated with a public cloud project.

## Example Usage

```terraform
data "ovh_cloud_project_database_database" "database" {
  service_name  = "XXX"
  engine        = "YYY"
  cluster_id    = "ZZZ"
  name          = "UUU"
}

output "database_name" {
  value = data.ovh_cloud_project_database_database.database.name
}
```

## Argument Reference

* `service_name` - (Required) The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.

* `engine` - (Required) The engine of the database cluster you want database information. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:
  * `mysql`
  * `postgresql`

* `cluster_id` - (Required) Cluster ID

* `name` - (Required) Name of the database.

## Attributes Reference

The following attributes are exported:

* `cluster_id` - See Argument Reference above.
* `created_at` - Date of the creation of the database.
* `default` - Defines if the database has been created by default.
* `engine` - See Argument Reference above.
* `id` - ID of the database.
* `name` - Name of the database.
* `service_name` - Current status of the database.
