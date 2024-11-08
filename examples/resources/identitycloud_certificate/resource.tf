resource "identitycloud_certificate" "example" {
  certificate = filebase64("mycert.pem")
  private_key = filebase64("key.pem")
}