// Code generated by ping-terraform-plugin-framework-generator

package resource

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	client "github.com/pingidentity/identitycloud-go-client/identitycloud"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/auth"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/providererror"
	internaltypes "github.com/pingidentity/terraform-provider-identitycloud/internal/types"
)

var (
	_ resource.Resource                = &cookieDomainsResource{}
	_ resource.ResourceWithConfigure   = &cookieDomainsResource{}
	_ resource.ResourceWithImportState = &cookieDomainsResource{}
)

func CookieDomainsResource() resource.Resource {
	return &cookieDomainsResource{}
}

type cookieDomainsResource struct {
	apiClient   *client.APIClient
	accessToken string
}

func (r *cookieDomainsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cookie_domains"
}

func (r *cookieDomainsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	resourceConfig, ok := req.ProviderData.(internaltypes.ResourceConfiguration)
	if !ok {
		resp.Diagnostics.AddError(providererror.InternalProviderError, "Invalid ProviderData when configuring resource. Please report this error in the provider's issue tracker.")
		return
	}
	r.apiClient = resourceConfig.ApiClient
	r.accessToken = resourceConfig.AccessToken
}

type cookieDomainsResourceModel struct {
	Domains types.List `tfsdk:"domains"`
}

func (r *cookieDomainsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	// domains default
	domainsDefault, diags := types.ListValue(types.StringType, nil)
	resp.Diagnostics.Append(diags...)

	resp.Schema = schema.Schema{
		Description: "Resource to create and manage the cookie domains.",
		Attributes: map[string]schema.Attribute{
			"domains": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The cookie domains. Defaults to an empty list.",
				Default:     listdefault.StaticValue(domainsDefault),
			},
		},
	}
}

func (model *cookieDomainsResourceModel) buildClientStruct() (*client.CookieDomains, diag.Diagnostics) {
	result := &client.CookieDomains{}
	// domains
	if !model.Domains.IsNull() {
		result.Domains = []string{}
		for _, domainsElement := range model.Domains.Elements() {
			result.Domains = append(result.Domains, domainsElement.(types.String).ValueString())
		}
	}

	return result, nil
}

// Build a default client struct to reset the resource to its default state
// If necessary, update this function to set any other values that should be present in the default state of the resource
func (model *cookieDomainsResource) buildDefaultClientStruct() *client.CookieDomains {
	result := &client.CookieDomains{}
	return result
}

func (state *cookieDomainsResourceModel) readClientResponse(response *client.CookieDomains) diag.Diagnostics {
	var respDiags, diags diag.Diagnostics
	// domains
	state.Domains, diags = types.ListValueFrom(context.Background(), types.StringType, response.Domains)
	respDiags.Append(diags...)
	return respDiags
}

// Set all non-primitive attributes to null with appropriate attribute types
func (r *cookieDomainsResource) emptyModel() cookieDomainsResourceModel {
	var model cookieDomainsResourceModel
	model.Domains = types.ListNull(types.StringType)
	return model
}

func (r *cookieDomainsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data cookieDomainsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic, since this is a singleton resource
	clientData, diags := data.buildClientStruct()
	resp.Diagnostics.Append(diags...)
	apiUpdateRequest := r.apiClient.CookieDomainsAPI.SetCookieDomains(auth.AuthContext(ctx, r.accessToken))
	apiUpdateRequest = apiUpdateRequest.Body(*clientData)
	responseData, _, err := r.apiClient.CookieDomainsAPI.SetCookieDomainsExecute(apiUpdateRequest)
	if err != nil {
		resp.Diagnostics.AddError("An error occurred while creating the cookieDomains", err.Error())
		return
	}

	// Read response into the model
	resp.Diagnostics.Append(data.readClientResponse(responseData)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *cookieDomainsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data cookieDomainsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	responseData, httpResp, err := r.apiClient.CookieDomainsAPI.GetCookieDomains(auth.AuthContext(ctx, r.accessToken)).Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			providererror.AddResourceNotFoundWarning(ctx, &resp.Diagnostics, "cookieDomains", httpResp)
			resp.State.RemoveResource(ctx)
		} else {
			providererror.ReportHttpError(ctx, &resp.Diagnostics, "An error occurred while reading the cookieDomains", err, httpResp)
		}
		return
	}

	// Read response into the model
	resp.Diagnostics.Append(data.readClientResponse(responseData)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *cookieDomainsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data cookieDomainsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic
	clientData, diags := data.buildClientStruct()
	resp.Diagnostics.Append(diags...)
	apiUpdateRequest := r.apiClient.CookieDomainsAPI.SetCookieDomains(auth.AuthContext(ctx, r.accessToken))
	apiUpdateRequest = apiUpdateRequest.Body(*clientData)
	responseData, httpResp, err := r.apiClient.CookieDomainsAPI.SetCookieDomainsExecute(apiUpdateRequest)
	if err != nil {
		providererror.ReportHttpError(ctx, &resp.Diagnostics, "An error occurred while updating the cookieDomains", err, httpResp)
		return
	}

	// Read response into the model
	resp.Diagnostics.Append(data.readClientResponse(responseData)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *cookieDomainsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// This resource is singleton, so it can't be deleted from the service.
	// Instead this delete method will attempt to set the resource to its default state on the service. If this isn't possible,
	// this method can be replaced with a no-op with a diagnostic warning message about being unable to set to the default state.
	// Update API call logic to reset to default
	defaultClientData := r.buildDefaultClientStruct()
	apiUpdateRequest := r.apiClient.CookieDomainsAPI.SetCookieDomains(auth.AuthContext(ctx, r.accessToken))
	apiUpdateRequest = apiUpdateRequest.Body(*defaultClientData)
	_, _, err := r.apiClient.CookieDomainsAPI.SetCookieDomainsExecute(apiUpdateRequest)
	if err != nil {
		resp.Diagnostics.AddError("An error occurred while resetting the cookieDomains", err.Error())
	}
}

func (r *cookieDomainsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// This resource has no identifier attributes, so the value passed in here doesn't matter. Just return an empty state struct.
	emptyState := r.emptyModel()
	resp.Diagnostics.Append(resp.State.Set(ctx, &emptyState)...)
}