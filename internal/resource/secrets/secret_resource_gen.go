// Copyright © 2025 Ping Identity Corporation

// Code generated by ping-terraform-plugin-framework-generator

package secrets

import (
	"context"
	"regexp"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	client "github.com/pingidentity/identitycloud-go-client/identitycloud"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/auth"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/providererror"
	internaltypes "github.com/pingidentity/terraform-provider-identitycloud/internal/types"
)

var (
	_ resource.Resource                = &secretResource{}
	_ resource.ResourceWithConfigure   = &secretResource{}
	_ resource.ResourceWithImportState = &secretResource{}
)

func SecretResource() resource.Resource {
	return &secretResource{}
}

type secretResource struct {
	apiClient                 *client.APIClient
	accessToken               *string
	serviceAccountTokenSource *client.ServiceAccountTokenSource
}

func (r *secretResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_secret"
}

func (r *secretResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
	r.serviceAccountTokenSource = resourceConfig.ServiceAccountConfig
}

type secretResourceModel struct {
	ActiveVersion     types.String `tfsdk:"active_version"`
	Description       types.String `tfsdk:"description"`
	Encoding          types.String `tfsdk:"encoding"`
	Id                types.String `tfsdk:"id"`
	LastChangeDate    types.String `tfsdk:"last_change_date"`
	LastChangedBy     types.String `tfsdk:"last_changed_by"`
	Loaded            types.Bool   `tfsdk:"loaded"`
	LoadedVersion     types.String `tfsdk:"loaded_version"`
	SecretId          types.String `tfsdk:"secret_id"`
	UseInPlaceholders types.Bool   `tfsdk:"use_in_placeholders"`
	ValueBase64       types.String `tfsdk:"value_base64"`
}

func (r *secretResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Resource to create and manage a secret.",
		Attributes: map[string]schema.Attribute{
			"active_version": schema.StringAttribute{
				Description: "Active version of the secret.",
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the secret.",
				Optional:    true,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(1000),
				},
				Default: stringdefault.StaticString(""),
			},
			"encoding": schema.StringAttribute{
				Required:    true,
				Description: "Type of base64 encoding used by the secret. Changing this value requires replacement of the resource. Supported values are `generic`, `pem`, `base64hmac`, `base64aes`.",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"generic",
						"pem",
						"base64hmac",
						"base64aes",
					),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"id": schema.StringAttribute{
				Description: "ID of the secret.",
				Computed:    true,
			},
			"last_change_date": schema.StringAttribute{
				Description: "Date of the last change to the secret.",
				Computed:    true,
			},
			"last_changed_by": schema.StringAttribute{
				Description: "User who last changed the secret.",
				Computed:    true,
			},
			"loaded": schema.BoolAttribute{
				Description: "Whether the secret is loaded.",
				Computed:    true,
			},
			"loaded_version": schema.StringAttribute{
				Description: "Version of the secret that is loaded.",
				Computed:    true,
			},
			"secret_id": schema.StringAttribute{
				Required:    true,
				Description: "ID of the secret. Must match the regex pattern `^esv-[a-z0-9_-]{1,124}$`.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^esv-[a-z0-9_-]{1,124}$"), ""),
				},
			},
			"use_in_placeholders": schema.BoolAttribute{
				Description: "Whether the secret is used in placeholders. Changing this value requires replacement of the resource.",
				Required:    true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"value_base64": schema.StringAttribute{
				Description: "Base64 encoded value of the secret. If you wish to change this value, use the `identitycloud_secret_version` resource to create a new version of this secret. Otherwise, changing this value will require replacement of the resource.",
				Required:    true,
				Sensitive:   true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile("^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$"), ""),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (model *secretResourceModel) buildClientStruct() (*client.EsvSecretCreateRequest, diag.Diagnostics) {
	result := &client.EsvSecretCreateRequest{}
	// description
	result.Description = model.Description.ValueStringPointer()
	// encoding
	result.Encoding = model.Encoding.ValueString()
	// use_in_placeholders
	result.UseInPlaceholders = model.UseInPlaceholders.ValueBool()
	// value_base64
	result.ValueBase64 = model.ValueBase64.ValueString()
	return result, nil
}

func (model *secretResourceModel) buildUpdateClientStruct() (*client.EsvSetDescriptionRequest, diag.Diagnostics) {
	result := &client.EsvSetDescriptionRequest{}
	// description
	result.Description = model.Description.ValueString()
	return result, nil
}

func (state *secretResourceModel) readClientResponse(response *client.EsvSecretResponse) diag.Diagnostics {
	// active_version
	state.ActiveVersion = types.StringValue(response.ActiveVersion)
	// description
	state.Description = types.StringValue(response.Description)
	// encoding
	state.Encoding = types.StringValue(response.Encoding)
	// id
	state.Id = types.StringValue(response.Id)
	// last_change_date
	state.LastChangeDate = types.StringValue(response.LastChangeDate.Format(time.RFC3339))
	// last_changed_by
	state.LastChangedBy = types.StringValue(response.LastChangedBy)
	// loaded
	state.Loaded = types.BoolValue(response.Loaded)
	// loaded_version
	state.LoadedVersion = types.StringValue(response.LoadedVersion)
	// use_in_placeholders
	state.UseInPlaceholders = types.BoolValue(response.UseInPlaceholders)
	return nil
}

func (r *secretResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data secretResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic
	clientData, diags := data.buildClientStruct()
	resp.Diagnostics.Append(diags...)
	apiCreateRequest := r.apiClient.SecretsAPI.CreateSecret(auth.AuthContext(ctx, r.accessToken, r.serviceAccountTokenSource), data.SecretId.ValueString())
	apiCreateRequest = apiCreateRequest.Body(*clientData)
	responseData, httpResp, err := r.apiClient.SecretsAPI.CreateSecretExecute(apiCreateRequest)
	if err != nil {
		providererror.ReportHttpError(ctx, &resp.Diagnostics, "An error occurred while creating the secret", err, httpResp)
		return
	}

	// Read response into the model
	resp.Diagnostics.Append(data.readClientResponse(responseData)...)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *secretResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data secretResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	responseData, httpResp, err := r.apiClient.SecretsAPI.GetSecret(auth.AuthContext(ctx, r.accessToken, r.serviceAccountTokenSource), data.SecretId.ValueString()).Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			providererror.AddResourceNotFoundWarning(ctx, &resp.Diagnostics, "secret", httpResp)
			resp.State.RemoveResource(ctx)
		} else {
			providererror.ReportHttpError(ctx, &resp.Diagnostics, "An error occurred while reading the secret", err, httpResp)
		}
		return
	}

	// Read response into the model
	resp.Diagnostics.Append(data.readClientResponse(responseData)...)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *secretResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state secretResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if !data.Description.Equal(state.Description) {
		// Update description API call logic
		clientData, diags := data.buildUpdateClientStruct()
		resp.Diagnostics.Append(diags...)
		apiUpdateRequest := r.apiClient.SecretsAPI.ActionSecret(auth.AuthContext(ctx, r.accessToken, r.serviceAccountTokenSource), data.SecretId.ValueString())
		apiUpdateRequest = apiUpdateRequest.Body(*clientData)
		apiUpdateRequest = apiUpdateRequest.Action("setDescription")
		httpResp, err := r.apiClient.SecretsAPI.ActionSecretExecute(apiUpdateRequest)
		if err != nil {
			providererror.ReportHttpError(ctx, &resp.Diagnostics, "An error occurred while updating the secret", err, httpResp)
			return
		}
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *secretResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data secretResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete API call logic
	_, httpResp, err := r.apiClient.SecretsAPI.DeleteSecret(auth.AuthContext(ctx, r.accessToken, r.serviceAccountTokenSource), data.SecretId.ValueString()).Execute()
	if err != nil && (httpResp == nil || httpResp.StatusCode != 404) {
		providererror.ReportHttpError(ctx, &resp.Diagnostics, "An error occurred while deleting the secret", err, httpResp)
	}
}

func (r *secretResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to secret_id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("secret_id"), req, resp)
}
