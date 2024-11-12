resource "identitycloud_certificate" "example" {
  certificate = file("mycert.pem")
  private_key = file("mykey.pem")
}