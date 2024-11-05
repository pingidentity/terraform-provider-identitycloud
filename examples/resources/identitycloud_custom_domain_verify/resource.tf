resource "identitycloud_custom_domain_verify" "example" {
  name = "mydomain.example.com"
  timeouts = {
    create = "30m"
  }
}