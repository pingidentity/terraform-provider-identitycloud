resource "identitycloud_custom_domain_verify" "example" {
  name = "auth.bxretail.org"
  timeouts = {
    create = "30m"
  }
}
