package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	client "github.com/pingidentity/identitycloud-go-client/identitycloud"
	internaltypes "github.com/pingidentity/terraform-provider-identitycloud/internal/types"
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

// Metadata returns the provider type name.
func (p *identityCloudProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "identitycloud"
}

// GetSchema defines the provider-level schema for configuration data.
// Schema defines the provider-level schema for configuration data.
func (p *identityCloudProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *identityCloudProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var resourceConfig internaltypes.ResourceConfiguration
	clientConfig := client.NewConfiguration()
	httpClient := &http.Client{}
	clientConfig.HTTPClient = httpClient
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
