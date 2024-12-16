# Contributing to the PingOne Advanced Identity Cloud Terraform Provider

We appreciate your help!  We welcome contributions in the form of creating issues or pull requests.

Know that:

1. If you have any questions, please ask!  We'll help as best we can.
2. While we appreciate perfect PRs, they are not essential. We will fix up any housekeeping changes before merging code.  If we find a PR that needs further work, we will guide you in the right direction.
3. We may not be able to respond quickly; our development cycles are on a priority basis.
4. We base our priorities on customer need and the number of votes on issues/PRs by the number of 👍 reactions.  If there is an existing issue or PR for something you would like, please vote!

## Creating Issues

Issues might include functional defects, unclear documentation or missing examples.  When creating issues, please follow the Bug report template.  Including as much information as possible helps others triage the issue and ultimately resolve the issue quicker.  Issues can be created through the "Issues" tab, or through [this link](https://github.com/pingidentity/terraform-provider-identitycloud/issues/new?assignees=&labels=bug&projects=&template=bug_report.md&title=).

## Code/Configuration Contribution

Resources and data sources are automatically generated and maintained from a Ping internal code generator. If your code contributions conflict with the code generator project, we may need to port the contributions to the internal generator first and relevant code re-generated.  The following are examples of contributions we can accept:

- Enhancing or fixing issues in core provider code (neither resource code, nor data source code).
- Enhancing or fixing issues in resources or data sources logic (including acceptance tests).  *These contributions are likely to be ported to the internal generator and the community PR closed without merge*
- Adding new fields to a resource/data source schema, or deprecating existing fields in a resource/data source schema.  *These contributions are likely to be ported to the internal generator and the community PR closed without merge*
- Enhancing and updating code quality CI checks and repository housekeeping assets.