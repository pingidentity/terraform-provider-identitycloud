provider "identitycloud" {
  tenant_environment_fqdn     = var.my_aic_tenant_environment_fqdn
  service_account_id          = var.my_aic_service_account_id
  service_account_private_key = file("private-key.jwk")
}
