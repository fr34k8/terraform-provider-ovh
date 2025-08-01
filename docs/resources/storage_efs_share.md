---
subcategory: "Enterprise File Storage"
---

# ovh_storage_efs_share (Resource)

Creates a share for an EFS service.

## Example Usage

```terraform
data "ovh_storage_efs" "efs" {
  service_name = "XXX"
}

resource "ovh_storage_efs_share" "share" {
  service_name = data.ovh_storage_efs.efs.service_name
  name         = "share"
  description  = "My share"
  protocol     = "NFS"
  size         = 100
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `protocol` (String) Share protocol
- `service_name` (String) Service name
- `size` (Number) Share size in Gigabytes

### Optional

- `description` (String) Share description
- `mount_point_name` (String) User-defined name used to generate human readable access path for the share
- `name` (String) Share name
- `snapshot_id` (String) Snapshot ID used to create the share

### Read-Only

- `created_at` (String) Share creation date
- `id` (String) Share ID
- `status` (String) Share status
