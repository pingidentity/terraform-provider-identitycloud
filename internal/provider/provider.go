// Copyright © 2025 Ping Identity Corporation

package provider

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	client "github.com/pingidentity/identitycloud-go-client/identitycloud"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/providererror"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/resource/certificate"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/resource/contentsecuritypolicy"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/resource/cookiedomains"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/resource/csrs"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/resource/customdomains"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/resource/promotion"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/resource/secrets"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/resource/ssocookie"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/resource/variable"
	internaltypes "github.com/pingidentity/terraform-provider-identitycloud/internal/types"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/utils"
)

// Ensure the implementation satisfies the expected interfacesß
var (
	_ provider.Provider = &identityCloudProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func NewFactory(version string) func() provider.Provider {
	return func() provider.Provider {
		return &identityCloudProvider{
			version: version,
		}
	}
}

// NewTestProvider is a helper function to simplify testing implementation.
func NewTestProvider() provider.Provider {
	return NewFactory("test")()
}

// identityCloudProvider is the provider implementation.
type identityCloudProvider struct {
	version string
}

// identityCloudProviderModel maps provider schema data to a Go type.
type identityCloudProviderModel struct {
	TenantEnvironmentFqdn    types.String `tfsdk:"tenant_environment_fqdn"`
	AccessToken              types.String `tfsdk:"access_token"`
	ServiceAccountId         types.String `tfsdk:"service_account_id"`
	ServiceAccountPrivateKey types.String `tfsdk:"service_account_private_key"`
}

// Metadata returns the provider type name.
func (p *identityCloudProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "identitycloud"
}

// GetSchema defines the provider-level schema for configuration data.
// Schema defines the provider-level schema for configuration data.
func (p *identityCloudProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"tenant_environment_fqdn": schema.StringAttribute{
				MarkdownDescription: "The fully qualified domain name of the tenant environment. Default value can be set with the `PINGAIC_TF_TENANT_ENV_FQDN` environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"access_token": schema.StringAttribute{
				MarkdownDescription: "Access token for the PingOne Advanced Identity Cloud Rest API. Default value can be set with the `PINGAIC_TF_ACCESS_TOKEN` environment variable. If `access_token` is provided, `service_account_id` and `service_account_private_key` should not be provided.",
				Optional:            true,
				Sensitive:           true,
			},
			"service_account_id": schema.StringAttribute{
				MarkdownDescription: "Service account ID for the PingOne Advanced Identity Cloud Rest API. The service account must have the following scopes: `fr:idc:certificate:*`, `fr:idc:content-security-policy:*`, `fr:idc:cookie-domain:*`, `fr:idc:custom-domain:*`, `fr:idc:esv:*`, `fr:idc:promotion:*`, `fr:idc:sso-cookie:*`. Default value can be set with the `PINGAIC_TF_SERVICE_ACCOUNT_ID` environment variable. If `service_account_id` and `service_account_private_key` are provided, `access_token` should not be provided.",
				Optional:            true,
				Sensitive:           true,
			},
			"service_account_private_key": schema.StringAttribute{
				MarkdownDescription: "Service account private key for the PingOne Advanced Identity Cloud Rest API. Default value can be set with the `PINGAIC_TF_SERVICE_ACCOUNT_PRIVATE_KEY` environment variable. If `service_account_id` and `service_account_private_key` are provided, `access_token` should not be provided.",
				Optional:            true,
				Sensitive:           true,
			},
		},
	}
}

func (p *identityCloudProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config *identityCloudProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide a tenant env FQDN to the provider
	var envFqdn string
	if !config.TenantEnvironmentFqdn.IsUnknown() && !config.TenantEnvironmentFqdn.IsNull() {
		envFqdn = config.TenantEnvironmentFqdn.ValueString()
	} else {
		envFqdn = os.Getenv("PINGAIC_TF_TENANT_ENV_FQDN")
	}

	if envFqdn == "" {
		resp.Diagnostics.AddAttributeError(path.Root("tenant_environment_fqdn"),
			providererror.InvalidProviderConfigurationError,
			"tenant_environment_fqdn provider attribute is required. If not set in the provider configuration, it can be set with the `PINGAIC_TF_TENANT_ENV_FQDN` environment variable.")
	} else {
		// Ensure the the FQDN can be parsed by url.Parse
		_, err := url.Parse(envFqdn)
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("tenant_environment_fqdn"),
				providererror.InvalidProviderConfigurationError,
				fmt.Sprintf("Invalid URL Format for FQDN '%s': %s", envFqdn, err.Error()),
			)
		}
	}

	// User may provide an access token to the provider
	var accessToken *string
	if !config.AccessToken.IsUnknown() && !config.AccessToken.IsNull() {
		accessToken = config.AccessToken.ValueStringPointer()
	} else {
		envToken := os.Getenv("PINGAIC_TF_ACCESS_TOKEN")
		if envToken != "" {
			accessToken = &envToken
		}
	}

	// Or users may provider a service account id and private key
	var serviceAccountId *string
	if !config.ServiceAccountId.IsUnknown() && !config.ServiceAccountId.IsNull() {
		serviceAccountId = config.ServiceAccountId.ValueStringPointer()
	} else {
		envId := os.Getenv("PINGAIC_TF_SERVICE_ACCOUNT_ID")
		if envId != "" {
			serviceAccountId = &envId
		}
	}

	var serviceAccountPrivateKey *string
	if !config.ServiceAccountPrivateKey.IsUnknown() && !config.ServiceAccountPrivateKey.IsNull() {
		serviceAccountPrivateKey = config.ServiceAccountPrivateKey.ValueStringPointer()
	} else {
		envKey := os.Getenv("PINGAIC_TF_SERVICE_ACCOUNT_PRIVATE_KEY")
		if envKey != "" {
			serviceAccountPrivateKey = &envKey
		}
	}

	// Ensure either a specific access token or the two service account attributes are provided
	if accessToken == nil && (serviceAccountId == nil || serviceAccountPrivateKey == nil) {
		resp.Diagnostics.AddError(providererror.InvalidProviderConfigurationError,
			"Either `access_token` or both `service_account_id` and `service_account_private_key` must be provided. If not set in the provider configuration, they can be set with the `PINGAIC_TF_ACCESS_TOKEN`, `PINGAIC_TF_SERVICE_ACCOUNT_ID`, and `PINGAIC_TF_SERVICE_ACCOUNT_PRIVATE_KEY` environment variables, respectively.")
	}

	if accessToken != nil && (serviceAccountId != nil || serviceAccountPrivateKey != nil) {
		resp.Diagnostics.AddError(providererror.InvalidProviderConfigurationError,
			"`access_token` should not be provided with either `service_account_id` or `service_account_private_key`. If not set in the provider configuration, they can be set with the `PINGAIC_TF_ACCESS_TOKEN`, `PINGAIC_TF_SERVICE_ACCOUNT_ID`, and `PINGAIC_TF_SERVICE_ACCOUNT_PRIVATE_KEY` environment variables, respectively.")
	}

	if resp.Diagnostics.HasError() {
		return
	}

	var serviceAccountTokenSource *client.ServiceAccountTokenSource
	if serviceAccountId != nil && serviceAccountPrivateKey != nil {
		serviceAccountTokenSource = &client.ServiceAccountTokenSource{
			TenantFqdn:               envFqdn,
			ServiceAccountId:         *serviceAccountId,
			ServiceAccountPrivateKey: *serviceAccountPrivateKey,
			Scopes: []string{
				"fr:idc:certificate:*",
				"fr:idc:content-security-policy:*",
				"fr:idc:cookie-domain:*",
				"fr:idc:custom-domain:*",
				"fr:idc:esv:*",
				"fr:idc:promotion:*",
				"fr:idc:sso-cookie:*",
			},
		}
	}

	resourceConfig := internaltypes.ResourceConfiguration{
		AccessToken:          accessToken,
		ServiceAccountConfig: serviceAccountTokenSource,
	}
	clientConfig := client.NewConfiguration()
	clientConfig.Servers = client.ServerConfigurations{
		{
			URL: fmt.Sprintf("https://%s", envFqdn),
		},
	}
	httpClient := &http.Client{}
	clientConfig.HTTPClient = httpClient
	userAgentSuffix := fmt.Sprintf("terraform-provider-identitycloud/%s", p.version)
	// The extra suffix for the user-agent is optional and is not considered a provider parameter.
	// We just use it directly from the environment variable, if set.
	userAgentExtraSuffix := os.Getenv("PINGAIC_TF_APPEND_USER_AGENT")
	if userAgentExtraSuffix != "" {
		userAgentSuffix += fmt.Sprintf(" %s", userAgentExtraSuffix)
	}
	clientConfig.UserAgentSuffix = utils.Pointer(userAgentSuffix)
	resourceConfig.ApiClient = client.NewAPIClient(clientConfig)
	resp.ResourceData = resourceConfig
	resp.DataSourceData = resourceConfig
	tflog.Info(ctx, "Configured identity cloud client", map[string]interface{}{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *identityCloudProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		promotion.PromotionLockDataSource,
		secrets.SecretVersionsDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *identityCloudProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		certificate.CertificateResource,
		cookiedomains.CookieDomainsResource,
		contentsecuritypolicy.ContentSecurityPolicyEnforcedResource,
		contentsecuritypolicy.ContentSecurityPolicyReportOnlyResource,
		csrs.CertificateSigningRequestExportResource,
		csrs.CertificateSigningRequestResponseResource,
		customdomains.CustomDomainsResource,
		customdomains.CustomDomainVerifyResource,
		promotion.PromotionLockResource,
		secrets.SecretResource,
		secrets.SecretVersionResource,
		ssocookie.SsoCookieResource,
		variable.VariableResource,
	}
}
