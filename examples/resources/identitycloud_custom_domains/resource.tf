resource "identitycloud_custom_domains" "example" {
  realm   = "alpha"
  domains = ["auth.bxretail.org", "sso.bxretail.org"]
}
