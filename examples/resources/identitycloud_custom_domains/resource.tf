resource "identitycloud_custom_domains" "example" {
  realm   = "alpha"
  domains = ["mydomain.example.com", "mydomain2.example.com"]
}