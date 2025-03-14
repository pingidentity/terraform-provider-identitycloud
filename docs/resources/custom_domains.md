---
page_title: "identitycloud_custom_domains Resource - terraform-provider-identitycloud"
subcategory: ""
description: |-
  Resource to create and manage the custom domains.
---

# identitycloud_custom_domains (Resource)

Resource to create and manage the custom domains.

Any custom domains will be validated by AIC when set. CNAME record verification can also be deactivated by submitting a ticket. See [the documentation on custom domains](https://docs.pingidentity.com/pingoneaic/latest/realms/custom-domains.html) for more information.

## Example Usage

```terraform
resource "identitycloud_custom_domains" "example" {
  realm   = "alpha"
  domains = ["auth.bxretail.org", "sso.bxretail.org"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `domains` (Set of String) The custom domains. Defaults to an empty set.
- `realm` (String) Realm for the domain. Supported values are `alpha`, `bravo`.
- `timeouts` (Attributes) (see [below for nested schema](#nestedatt--timeouts))

<a id="nestedatt--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m", as a time to wait for DNS record changes to propagate for verification on initial create of the resource. Valid time units are "s" (seconds), "m" (minutes), "h" (hours). The default is 1 minute.
- `update` (String) A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m", as a time to wait for DNS record changes to propagate for verification on update of the resource. Valid time units are "s" (seconds), "m" (minutes), "h" (hours). The default is 1 minute.

## Import

Import is supported using the following syntax:

~> realm_id should be either `alpha` or `bravo`.

```shell
terraform import identitycloud_custom_domains.example realm_id
```