---
page_title: "identitycloud_secret Resource - terraform-provider-identitycloud"
subcategory: ""
description: |-
  Resource to create and manage a secret.
---

# identitycloud_secret (Resource)

Resource to create and manage a secret.

The statuses of individual versions of a secret should be managed with the `identitycloud_secret_version` resource.

## Example Usage

```terraform
resource "identitycloud_secret" "example" {
  secret_id           = "esv-examplesecret"
  description         = "my example secret"
  encoding            = "generic"
  use_in_placeholders = false
  value_base64        = base64encode(var.example_secret_value)
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `encoding` (String) Type of base64 encoding used by the secret. Changing this value requires replacement of the resource. Supported values are `generic`, `pem`, `base64hmac`, `base64aes`.
- `secret_id` (String) ID of the secret. Must match the regex pattern `^esv-[a-z0-9_-]{1,124}$`.
- `use_in_placeholders` (Boolean) Whether the secret is used in placeholders. Changing this value requires replacement of the resource.
- `value_base64` (String, Sensitive) Base64 encoded value of the secret. Changing this value will create a new version of the secret.

### Optional

- `description` (String) Description of the secret.

### Read-Only

- `active_version` (String) Active version of the secret.
- `id` (String) ID of the secret.
- `last_change_date` (String) Date of the last change to the secret.
- `last_changed_by` (String) User who last changed the secret.
- `loaded` (Boolean) Whether the secret is loaded.
- `loaded_version` (String) Version of the secret that is loaded.

## Import

Import is supported using the following syntax:

```shell
terraform import identitycloud_secret.example secret_id
```