---
subcategory : "Managed Databases"
---

# ovh_cloud_project_database_kafka_acl

Creates an ACL for a kafka cluster associated with a public cloud project.

## Example Usage

```terraform
data "ovh_cloud_project_database" "kafka" {
  service_name  = "XXX"
  engine        = "kafka"
  id            = "ZZZ"
}

resource "ovh_cloud_project_database_kafka_acl" "acl" {
  service_name    = data.ovh_cloud_project_database.kafka.service_name
  cluster_id      = data.ovh_cloud_project_database.kafka.id
  permission      = "read"
  topic           = "mytopic"
  username        = "johndoe"
}
```

## Argument Reference

The following arguments are supported:

* `service_name` - (Required, Forces new resource) The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.

* `cluster_id` - (Required, Forces new resource) Cluster ID.

* `permission` - (Required, Forces new resource) Permission to give to this username on this topic. Available permissions:
  * `admin`
  * `read`
  * `write`
  * `readwrite`

* `topic` - (Required, Forces new resource) Topic affected by this ACL.

* `username` - (Required, Forces new resource) Username affected by this ACL.

## Attributes Reference

The following attributes are exported:

* `cluster_id` - See Argument Reference above.
* `id` - ID of the ACL.
* `permission` - See Argument Reference above.
* `service_name` - See Argument Reference above.
* `topic` - See Argument Reference above.
* `username` - See Argument Reference above.

## Timeouts

```terraform
resource "ovh_cloud_project_database_kafka_acl" "acl" {
  # ...

  timeouts {
    create = "1h"
    delete = "45m"
  }
}
```
* `create` - (Default 20m)
* `delete` - (Default 20m)

## Import

OVHcloud Managed kafka clusters ACLs can be imported using the `service_name`, `cluster_id` and `id` of the acl, separated by "/" E.g.,

```bash
$ terraform import ovh_cloud_project_database_kafka_acl.my_acl service_name/cluster_id/id
```
