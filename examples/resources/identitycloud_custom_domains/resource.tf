resource "identitycloud_custom_domains" "example" {
  realm   = "alpha"
  domains = ["mydomain1", "mydomain2"]
}