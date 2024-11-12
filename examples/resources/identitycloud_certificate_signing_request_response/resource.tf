resource "identitycloud_certificate_signing_request_export" "example" {
  common_name  = "Test CN"
  organization = "Test Org"
  country      = "US"
}

resource "tls_locally_signed_cert" "example" {
  cert_request_pem   = identitycloud_certificate_signing_request_export.example.request
  ca_private_key_pem = file("ca_private_key.pem")
  ca_cert_pem        = file("ca_cert.pem")

  validity_period_hours = 12

  allowed_uses = [
    "server_auth",
  ]
}

resource "identitycloud_certificate_signing_request_response" "example" {
  certificate                    = tls_locally_signed_cert.example.cert_pem
  certificate_signing_request_id = identitycloud_certificate_signing_request_export.example.id
}