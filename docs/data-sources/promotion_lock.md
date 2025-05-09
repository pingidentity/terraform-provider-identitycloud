---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "identitycloud_promotion_lock Data Source - terraform-provider-identitycloud"
subcategory: ""
description: |-
  Resource to create and manage the promotion lock process.
---

# identitycloud_promotion_lock (Data Source)

Resource to create and manage the promotion lock process.

## Example Usage

```terraform
data "identitycloud_promotion_lock" "example" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `description` (String) Description of the state of the lock.
- `id` (String) Promotion unique identifier.
- `lower_env` (Attributes) Lower environment lock status. (see [below for nested schema](#nestedatt--lower_env))
- `promotion_id` (String) Promotion unique identifier.
- `result` (String) The lock status of the environment. Supported values are `unlocking`, `unlocked`, `locking`, `locked`, `error`.
- `upper_env` (Attributes) Upper environment lock status. (see [below for nested schema](#nestedatt--upper_env))

<a id="nestedatt--lower_env"></a>
### Nested Schema for `lower_env`

Read-Only:

- `promotion_id` (String) Promotion unique identifier.
- `proxy_state` (String) Proxy state of the lock.
- `state` (String) State of the lock.


<a id="nestedatt--upper_env"></a>
### Nested Schema for `upper_env`

Read-Only:

- `promotion_id` (String) Promotion unique identifier.
- `proxy_state` (String) Proxy state of the lock.
- `state` (String) State of the lock.
