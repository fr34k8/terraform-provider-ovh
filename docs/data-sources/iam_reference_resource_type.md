---
subcategory : "Account Management (IAM)"
---

# ovh_iam_reference_resource_type (Data Source)

Use this data source to list all the IAM resource types.

## Example Usage

```terraform
data "ovh_iam_reference_resource_type" "types" {
}
```

## Argument Reference

## Attributes Reference

* `id` - hash of the list of the resource types
* `types` - List of resource types
