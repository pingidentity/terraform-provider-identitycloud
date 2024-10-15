package provider

import (
	"context"
	"fmt"
	"net/http"
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
	TenantEnvironmentFqdn types.String `tfsdk:"tenant_environment_fqdn"`
	AccessToken           types.String `tfsdk:"access_token"`
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
				MarkdownDescription: "Access token for the PingOne Advanced Identity Cloud Rest API. Default value can be set with the `PINGAIC_TF_ACCESS_TOKEN` environment variable.",
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
		resp.Diagnostics.AddAttributeError(path.Root("tenant_environment_fqdn"), providererror.InvalidProviderConfiguration, "tenant_environment_fqdn provider attribute is required. If not set in the provider configuration, it can be set with the `PINGAIC_TF_TENANT_ENV_FQDN` environment variable.")
	} else {
		//TODO validate the FQDN
	}
	// User must provide an access token to the provider
	var accessToken string
	if !config.AccessToken.IsUnknown() && !config.AccessToken.IsNull() {
		accessToken = config.AccessToken.ValueString()
	} else {
		accessToken = os.Getenv("PINGAIC_TF_ACCESS_TOKEN")
	}

	if accessToken == "" {
		resp.Diagnostics.AddAttributeError(path.Root("access_token"), providererror.InvalidProviderConfiguration, "access_token provider attribute is required. If not set in the provider configuration, it can be set with the `PINGAIC_TF_ACCESS_TOKEN` environment variable.")
	}

	resourceConfig := internaltypes.ResourceConfiguration{
		TenantEnvironmentFqdn: envFqdn,
		AccessToken:           accessToken,
	}
	clientConfig := client.NewConfiguration()
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
	return []func() datasource.DataSource{}
}

// Resources defines the resources implemented in the provider.
func (p *identityCloudProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
