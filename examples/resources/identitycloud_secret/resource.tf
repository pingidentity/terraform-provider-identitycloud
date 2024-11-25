resource "identitycloud_secret" "example" {
  secret_id           = "esv-examplesecret"
  description         = "my example secret"
  encoding            = "generic"
  use_in_placeholders = false
  value_base64        = base64encode(var.example_secret_value)
}