resource "identitycloud_certificate" "example" {
  certificate = file("/path/to/mycert.pem")
  private_key = file("/path/to/mykey.pem")
}
