resource "identitycloud_certificate" "example" {
  certificate = filebase64("mycert.pem")
  private_key = var.certificate_private_key
}