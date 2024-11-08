resource "identitycloud_secret" "example" {
  secret_id           = "esv-examplesecret"
  encoding            = "generic"
  use_in_placeholders = false
  value_base64        = base64encode(var.example_secret_value)
}

resource "identitycloud_secret_version" "example" {
  secret_id  = identitycloud_secret.example.secret_id
  version_id = identitycloud_secret.example.active_version
  status     = "ENABLED"
}