---
subcategory : "vRack"
---

# ovh_vrack_cloudproject

Attach a Public Cloud Project to a VRack.

## Example Usage

```terraform
resource "ovh_vrack_cloudproject" "vcp" {
  service_name = "12345"
  project_id   = "67890"
}
```

## Argument Reference

The following arguments are supported:

* `service_name` - (Required) The service name of the vrack. If omitted, the `OVH_VRACK_SERVICE` environment variable is used.

* `project_id` - (Required) The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.

## Attributes Reference

The following attributes are exported:

* `service_name` - See Argument Reference above.
* `project_id` - See Argument Reference above.

## Import

Attachment of a public cloud project and a VRack can be imported using the `service_name` (vRack identifier) and the `project_id` (Cloud Project identifier), separated by "/" E.g.,

```bash
$ terraform import ovh_vrack_cloudproject.myattach service_name/project_id
```
