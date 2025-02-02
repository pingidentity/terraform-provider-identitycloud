resource "identitycloud_variable" "example" {
  variable_id     = "esv-myvariable1"
  description     = "My variable"
  expression_type = "list"
  value_base64 = base64encode(jsonencode(
    [
      "value1",
      "value2",
      "value3"
    ]
  ))
}
