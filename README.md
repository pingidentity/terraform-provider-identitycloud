# PingOne Advanced Identity Cloud Terraform Provider
The PingOne Advanced Identity Cloud Terraform provider is a plugin for [Terraform](https://www.terraform.io/) that supports the management of Advanced Identity Cloud tenants. This provider is maintained internally by the Ping Identity team.

# Disclaimer - Provider in Development
The PingOne Advanced Identity Cloud Terraform provider is under active development. As such, consumers must have flexibility for breaking changes until the `1.0.0` release. When using the provider within an automated pipeline prior to `1.0.0`, it is recommended to pin the provider version similar to `version = "~> 0.5.0"` to avoid experiencing an unexpected pipeline failure as the result of a provider change. Enhancements, bug fixes, notes and breaking changes can be found on the [Changelog](CHANGELOG.md). If issues are found, please raise a [github issue](https://github.com/pingidentity/terraform-provider-identitycloud/issues/new?assignees=&labels=bug&projects=&template=bug_report.md&title=) on this project.

## Getting Started

Provider documentation can be found on the [Terraform registry](https://registry.terraform.io/providers/pingidentity/identitycloud/latest), and at [terraform.pingidentity.com](https://terraform.pingidentity.com).

## Requirements
* Terraform 1.4+
* Go 1.22.3+ (for local development builds)

## Useful Links
* [Discuss the PingOne Advanced Identity Cloud Terraform Provider](https://support.pingidentity.com/s/topic/0TO1W000000IF30WAG/pingdevops)
* [Ping Identity Home](https://www.pingidentity.com/en.html)

Extended documentation can be found at:
* [PingOne Advanced Identity Cloud documentation](https://docs.pingidentity.com/pingoneaic/latest/)
* [Ping Identity Developer Portal](https://developer.pingidentity.com/en.html)

# Contributing
We appreciate your help! To contribute through logging issues or creating pull requests, please read the [contribution guidelines](CONTRIBUTING.md).
