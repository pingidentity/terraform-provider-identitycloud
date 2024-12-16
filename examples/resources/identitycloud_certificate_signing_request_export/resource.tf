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
