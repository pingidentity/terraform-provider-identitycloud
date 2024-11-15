package certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	client "github.com/pingidentity/identitycloud-go-client/identitycloud"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/auth"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/providererror"
)

// Only the active attribute can be updated after creation
func (r *certificateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data certificateResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic
	clientData := client.UpdateCertificateRequest{
		Active: data.Active.ValueBoolPointer(),
	}
	apiUpdateRequest := r.apiClient.CertificatesAPI.UpdateCertificateByID(auth.AuthContext(ctx, r.accessToken, r.serviceAccountTokenSource), data.Id.ValueString())
	apiUpdateRequest = apiUpdateRequest.Body(clientData)
	responseData, httpResp, err := r.apiClient.CertificatesAPI.UpdateCertificateByIDExecute(apiUpdateRequest)
	if err != nil {
		providererror.ReportHttpError(ctx, &resp.Diagnostics, "An error occurred while updating the certificate", err, httpResp)
		return
	}

	// Read response into the model
	resp.Diagnostics.Append(data.readClientResponse(responseData)...)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
