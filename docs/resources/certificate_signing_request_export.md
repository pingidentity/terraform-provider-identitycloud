---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "identitycloud_certificate_signing_request_export Resource - terraform-provider-identitycloud"
subcategory: ""
description: |-
  Resource to create and manage a certificate signing request.
---

# identitycloud_certificate_signing_request_export (Resource)

Resource to create and manage a certificate signing request.

## Example Usage

```terraform
resource "identitycloud_certificate_signing_request_export" "example" {
  algorithm                 = "rsa"
  business_category         = "Business Entity"
  city                      = "Austin"
  common_name               = "bxretail.org Demo"
  country                   = "US"
  email                     = "admin@bxretail.org"
  organization              = "BXRetail"
  organizational_unit       = "Auth Services"
  state                     = "TX"
  street_address            = "1234 Example St"
  subject_alternative_names = ["bxretail.org"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `algorithm` (String) The algorithm for the private key. The encryption algorithm will either be RSA-2048 or ECDSA P-256 depending on the algorithm choice. The default is `rsa`. Supported values are `rsa`, `ecdsa`.
- `business_category` (String) Category of business, such as "Private Organization", “Government Entity”, “Business Entity”, or “Non-Commercial Entity”. Relevant for EV certificates.
- `city` (String) City for the CSR
- `common_name` (String) Domain name that the SSL certificate is securing. At least one of `common_name` or `subject_alternative_names` must be specified.
- `country` (String) Two-letter ISO-3166 country code
- `email` (String) Email for the CSR
- `jurisdiction_city` (String) This field contains only information relevant to the Jurisdiction of Incorporation or Registration. Relevant for EV certificates.
- `jurisdiction_country` (String) This field contains only information relevant to the Jurisdiction of Incorporation or Registration. Relevant for EV certificates.
- `jurisdiction_state` (String) This field contains only information relevant to the Jurisdiction of Incorporation or Registration. Relevant for EV certificates.
- `organization` (String) Full name of company
- `organizational_unit` (String) Company section or department
- `postal_code` (String) Postal code for the CSR
- `serial_number` (String) The Registration (or similar) Number assigned to the Subject by the Incorporating or Registration Agency in its Jurisdiction of Incorporation or Registration. Relevant for EV certificates.
- `state` (String) State for the CSR
- `street_address` (String) Street address for the CSR
- `subject_alternative_names` (Set of String) Additional domain or domains that the SSL certificate is securing. At least one of `common_name` or `subject_alternative_names` must be specified.

### Read-Only

- `certificate_id` (String) The ID of the certificate created from this CSR if the CSR has been completed.
- `created_date` (String) The date the CSR was created
- `id` (String) The unique identifier for the CSR
- `request` (String) PEM formatted CSR.
- `subject` (String) The CSR subject
