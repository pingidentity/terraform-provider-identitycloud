resource "identitycloud_content_security_policy_report_only" "example" {
  active = true
  directives = {
    "property1" : ["value1"],
    "property2" : ["value2"],
  }
}
