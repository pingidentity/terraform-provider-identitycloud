// Code generated by ping-terraform-plugin-framework-generator

package contentsecuritypolicy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	client "github.com/pingidentity/identitycloud-go-client/identitycloud"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/auth"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/providererror"
	internaltypes "github.com/pingidentity/terraform-provider-identitycloud/internal/types"
)

var (
	_ resource.Resource                = &contentSecurityPolicyReportOnlyResource{}
	_ resource.ResourceWithConfigure   = &contentSecurityPolicyReportOnlyResource{}
	_ resource.ResourceWithImportState = &contentSecurityPolicyReportOnlyResource{}
)

func ContentSecurityPolicyReportOnlyResource() resource.Resource {
	return &contentSecurityPolicyReportOnlyResource{}
}

type contentSecurityPolicyReportOnlyResource struct {
	apiClient   *client.APIClient
	accessToken string
}

func (r *contentSecurityPolicyReportOnlyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_content_security_policy_report_only"
}

func (r *contentSecurityPolicyReportOnlyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

type contentSecurityPolicyReportOnlyResourceModel struct {
	Active     types.Bool `tfsdk:"active"`
	Directives types.Map  `tfsdk:"directives"`
}

func (r *contentSecurityPolicyReportOnlyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Resource to create and manage the report-only content security policy.",
		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Whether the policy is active. The default value is `false`.",
			},
			"directives": schema.MapAttribute{
				ElementType: types.ListType{
					ElemType: types.StringType,
				},
				Optional:    true,
				Computed:    true,
				Default:     mapdefault.StaticValue(types.MapValueMust(types.ListType{ElemType: types.StringType}, nil)),
				Description: "The directives to enforce.",
			},
		},
	}
}

func (model *contentSecurityPolicyReportOnlyResourceModel) buildClientStruct() (*client.ContentSecurityPolicy, diag.Diagnostics) {
	result := &client.ContentSecurityPolicy{}
	// active
	result.Active = model.Active.ValueBoolPointer()
	// directives
	if !model.Directives.IsNull() {
		result.Directives = &map[string][]string{}
		for key, directivesElement := range model.Directives.Elements() {
			directivesValue := []string{}
			for _, directivesInnerElement := range directivesElement.(types.List).Elements() {
				directivesValue = append(directivesValue, directivesInnerElement.(types.String).ValueString())
			}
			(*result.Directives)[key] = directivesValue
		}
	}

	return result, nil
}

// Build a default client struct to reset the resource to its default state
// If necessary, update this function to set any other values that should be present in the default state of the resource
func (model *contentSecurityPolicyReportOnlyResource) buildDefaultClientStruct() *client.ContentSecurityPolicy {
	result := &client.ContentSecurityPolicy{}
	return result
}

func (state *contentSecurityPolicyReportOnlyResourceModel) readClientResponse(response *client.ContentSecurityPolicy) diag.Diagnostics {
	var respDiags, diags diag.Diagnostics
	// active
	state.Active = types.BoolPointerValue(response.Active)
	// directives
	directivesElementType := types.ListType{ElemType: types.StringType}
	if response.Directives == nil {
		state.Directives = types.MapNull(directivesElementType)
	} else {
		state.Directives, diags = types.MapValueFrom(context.Background(), directivesElementType, (*response.Directives))
		respDiags.Append(diags...)
	}
	return respDiags
}

// Set all non-primitive attributes to null with appropriate attribute types
func (r *contentSecurityPolicyReportOnlyResource) emptyModel() contentSecurityPolicyReportOnlyResourceModel {
	var model contentSecurityPolicyReportOnlyResourceModel
	// directives
	directivesElementType := types.ListType{ElemType: types.StringType}
	model.Directives = types.MapNull(directivesElementType)
	return model
}

func (r *contentSecurityPolicyReportOnlyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data contentSecurityPolicyReportOnlyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic, since this is a singleton resource
	clientData, diags := data.buildClientStruct()
	resp.Diagnostics.Append(diags...)
	apiUpdateRequest := r.apiClient.ContentSecurityPolicyAPI.SetReportOnlyContentSecurityPolicy(auth.AuthContext(ctx, r.accessToken))
	apiUpdateRequest = apiUpdateRequest.Body(*clientData)
	responseData, _, err := r.apiClient.ContentSecurityPolicyAPI.SetReportOnlyContentSecurityPolicyExecute(apiUpdateRequest)
	if err != nil {
		resp.Diagnostics.AddError("An error occurred while creating the contentSecurityPolicyReportOnly", err.Error())
		return
	}

	// Read response into the model
	resp.Diagnostics.Append(data.readClientResponse(responseData)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *contentSecurityPolicyReportOnlyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data contentSecurityPolicyReportOnlyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	responseData, httpResp, err := r.apiClient.ContentSecurityPolicyAPI.GetReportOnlyContentSecurityPolicy(auth.AuthContext(ctx, r.accessToken)).Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			providererror.AddResourceNotFoundWarning(ctx, &resp.Diagnostics, "contentSecurityPolicyReportOnly", httpResp)
			resp.State.RemoveResource(ctx)
		} else {
			providererror.ReportHttpError(ctx, &resp.Diagnostics, "An error occurred while reading the contentSecurityPolicyReportOnly", err, httpResp)
		}
		return
	}

	// Read response into the model
	resp.Diagnostics.Append(data.readClientResponse(responseData)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *contentSecurityPolicyReportOnlyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data contentSecurityPolicyReportOnlyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic
	clientData, diags := data.buildClientStruct()
	resp.Diagnostics.Append(diags...)
	apiUpdateRequest := r.apiClient.ContentSecurityPolicyAPI.SetReportOnlyContentSecurityPolicy(auth.AuthContext(ctx, r.accessToken))
	apiUpdateRequest = apiUpdateRequest.Body(*clientData)
	responseData, httpResp, err := r.apiClient.ContentSecurityPolicyAPI.SetReportOnlyContentSecurityPolicyExecute(apiUpdateRequest)
	if err != nil {
		providererror.ReportHttpError(ctx, &resp.Diagnostics, "An error occurred while updating the contentSecurityPolicyReportOnly", err, httpResp)
		return
	}

	// Read response into the model
	resp.Diagnostics.Append(data.readClientResponse(responseData)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *contentSecurityPolicyReportOnlyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// This resource is singleton, so it can't be deleted from the service.
	// Instead this delete method will attempt to set the resource to its default state on the service. If this isn't possible,
	// this method can be replaced with a no-op with a diagnostic warning message about being unable to set to the default state.
	// Update API call logic to reset to default
	defaultClientData := r.buildDefaultClientStruct()
	apiUpdateRequest := r.apiClient.ContentSecurityPolicyAPI.SetReportOnlyContentSecurityPolicy(auth.AuthContext(ctx, r.accessToken))
	apiUpdateRequest = apiUpdateRequest.Body(*defaultClientData)
	_, _, err := r.apiClient.ContentSecurityPolicyAPI.SetReportOnlyContentSecurityPolicyExecute(apiUpdateRequest)
	if err != nil {
		resp.Diagnostics.AddError("An error occurred while resetting the contentSecurityPolicyReportOnly", err.Error())
	}
}

func (r *contentSecurityPolicyReportOnlyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// This resource has no identifier attributes, so the value passed in here doesn't matter. Just return an empty state struct.
	emptyState := r.emptyModel()
	resp.Diagnostics.Append(resp.State.Set(ctx, &emptyState)...)
}