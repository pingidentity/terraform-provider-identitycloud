resource "identitycloud_promotion_lock" "example" {
  retry_timeouts = {
    create = "10m"
    delete = "10m"
  }
}