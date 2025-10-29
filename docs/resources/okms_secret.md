
---
subcategory : "Key Management Service (KMS)"
---

# ovh_okms_secret (Resource)

Manages a secret stored in OVHcloud KMS.

> WARNING: `version.data` is marked **Sensitive** but still ends up in the state file. To mitigate that, it is recommended to protect your state with encryption and access controls. Avoid committing it to source control.

## Example Usage

Create a secret whose value is a JSON object. Use `jsonencode()` to produce a deterministic JSON string (ordering/whitespace) to minimize diffs.

```terraform
resource "ovh_okms_secret" "example" {
	okms_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	path    = "app/api_credentials"

	metadata = {
		max_versions             = 10         # keep last 10 versions
		cas_required             = true       # enforce optimistic concurrency control (server will require current secret version on the cas attribute to allow update)
		deactivate_version_after = "0s"       # keep versions active indefinitely (example)
		custom_metadata = {
			environment = "prod"
			owner       = "payments-team"
		}
	}

	# Initial version (will create version 1)
	version = {
		data = jsonencode({
			api_key    = var.api_key
			api_secret = var.api_secret
		})
	}
}

# Reading a field from the secret version data
locals {
	secret_json = jsondecode(ovh_okms_secret.example.version.data)
}

output "api_key" {
	value     = local.secret_json.api_key
	sensitive = true
}
```

Updating the secret (new version) while enforcing optimistic concurrency control using CAS:

```terraform
resource "ovh_okms_secret" "example" {
	okms_id = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	path    = "app/api_credentials"

	# Ensure no concurrent update happened: set cas to the current version
	# (metadata.current_version is populated after first apply)
	cas = ovh_okms_secret.example.metadata.current_version

	metadata = {
		cas_required = true
	}

	version = {
		data = jsonencode({
			api_key    = var.api_key
			api_secret = var.new_api_secret  # changed value -> creates new version
		})
	}
}
```

## Argument Reference

The following arguments are supported:

### Required

- `okms_id` (String) ID of the OKMS service to create the secret in.
- `path` (String) Secret path (acts as the secret identifier within the OKMS instance). Immutable after creation.
- `version` (Block) Definition of the version to create/update. See Version Block below. (On updates providing a new `version.data` creates a new version.)

### Optional

- `cas` (Number) Check‑and‑set parameter used only on update (if `cas_required` metadata is set to true) to enforce optimistic concurrency control: its value must equal the current secret version (`metadata.current_version`) for the update to succeed. Ignored on create.
- `metadata` (Block) Secret metadata configuration (subset of fields are user-settable). See Metadata Block below.

### Metadata Block

User configurable attributes inside `metadata`:

- `cas_required` (Boolean) If true, the server will enforce optimistic concurrency control by requiring the `cas` parameter to match the current version number on every write (update) request.
- `custom_metadata` (Map of String) Arbitrary key/value metadata.
- `deactivate_version_after` (String) Duration (e.g. `"24h"`) after which a version is deactivated. `"0s"` (default) means never automatically deactivate.
- `max_versions` (Number) Number of versions to retain (default 10). Older versions beyond this limit are pruned.

Computed (read‑only) metadata attributes:

- `created_at` (String) Creation timestamp of the secret.
- `updated_at` (String) Last update timestamp.
- `current_version` (Number) Current (latest) version number.
- `oldest_version` (Number) Oldest retained version number.

### Version Block

Required attribute:

- `data` (String, Sensitive) Secret payload. Commonly set with `jsonencode(...)` so that Terraform comparisons are stable. Any valid JSON (object, array, string, number, bool) is accepted. Changing `data` creates a new secret version.

Computed (read‑only) attributes:

- `id` (Number) Version number.
- `created_at` (String) Version creation timestamp.
- `deactivated_at` (String) Deactivation timestamp if the version was deactivated.
- `state` (String) Version state (e.g. `ACTIVE`).

## Attributes Reference (Read-Only)

In addition to arguments above, the following attributes are exported:

- `iam` (Block) IAM metadata: `display_name`, `id`, `tags`, `urn`.
- `metadata.*` computed fields as listed above.
- `version.*` computed fields as listed above.

## Behavior & Notes

- Updating with a new `version.data` performs an API PUT that creates a new version; the previous version remains (subject to `max_versions`).
- If `cas_required` is true, all write operations must include a correct `cas` query parameter (the expected current version number). Set `cas = ovh_okms_secret.example.metadata.current_version` to enforce optimistic concurrency. A mismatch causes the API to reject the update (preventing overwriting unseen changes).
- `cas` is ignored on create (no existing version).
