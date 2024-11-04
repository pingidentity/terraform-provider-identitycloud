resource "identitycloud_certificate_signing_request" "example" {
  algorithm                 = "rsa"
  business_category         = "Example"
  city                      = "Austin"
  common_name               = "Ping"
  country                   = "US"
  email                     = "example@example.com"
  jurisdiction_city         = "Austin"
  jurisdiction_country      = "US"
  jurisdiction_state        = "TX"
  organization              = "Ping"
  organizational_unit       = "Example"
  postal_code               = "78701"
  serial_number             = "123456"
  state                     = "TX"
  street_address            = "1234 Example St"
  subject_alternative_names = ["example.com"]
}