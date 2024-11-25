resource "identitycloud_secret" "example" {
  secret_id           = "esv-examplesecret"
  encoding            = "generic"
  use_in_placeholders = false
  value_base64        = base64encode(var.example_secret_value)
}

resource "identitycloud_secret_version" "version2" {
  secret_id    = identitycloud_secret.example.secret_id
  value_base64 = base64encode(var.example_version2_secret_value)
}

resource "identitycloud_secret_version" "version3" {
  depends_on   = [identitycloud_secret_version.version2]
  secret_id    = identitycloud_secret.example.secret_id
  value_base64 = base64encode(var.example_version3_secret_value)
}