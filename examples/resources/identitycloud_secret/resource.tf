resource "identitycloud_secret" "example" {
  variable_id  = "esv-mysecret1"
  description  = "My secret"
  encoding     = "generic"
  value_base64 = base64encode("secretvalue")
}