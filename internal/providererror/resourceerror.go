package providererror

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Report a 404 as a warning for resources
func AddResourceNotFoundWarning(ctx context.Context, diagnostics *diag.Diagnostics, resourceType string, httpResp *http.Response) {
	diagnostics.AddWarning("Resource not found", fmt.Sprintf("The requested %s resource configuration cannot be found in the Identity Cloud tenant.  If the requested resource is managed in Terraform's state, it may have been removed outside of Terraform.", resourceType))
	if httpResp != nil {
		body, err := io.ReadAll(httpResp.Body)
		if err == nil {
			tflog.Debug(ctx, "Error HTTP response body: "+string(body))
		} else {
			tflog.Warn(ctx, "Failed to read HTTP response body: "+err.Error())
		}
	}
}

type AicErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ReadErrorResponse(ctx context.Context, httpResp *http.Response) (*AicErrorResponse, []byte) {
	if httpResp == nil {
		return nil, nil
	}
	var aicError AicErrorResponse

	defer httpResp.Body.Close()
	body, err := io.ReadAll(httpResp.Body)
	if err == nil {
		err = json.Unmarshal(body, &aicError)
		if err == nil {
			return &aicError, body
		}
		return nil, body
	}
	return nil, nil
}

// Report an HTTP error
func ReportHttpError(ctx context.Context, diagnostics *diag.Diagnostics, errorSummary string, err error, httpResp *http.Response) {
	httpErrorPrinted := false
	var internalError error
	var body []byte
	if httpResp != nil {
		body, internalError = io.ReadAll(httpResp.Body)
		if internalError == nil {
			ReportHttpErrorBody(ctx, diagnostics, errorSummary, err, body)
			httpErrorPrinted = true
		}
	}
	if !httpErrorPrinted {
		if internalError != nil {
			tflog.Warn(ctx, "Failed to read HTTP response body: "+internalError.Error())
		}
		diagnostics.AddError(AicAPIError, errorSummary+"\n"+err.Error())
	}
}

// Report an HTTP error
func ReportHttpErrorBody(ctx context.Context, diagnostics *diag.Diagnostics, errorSummary string, err error, httpRespBody []byte) {
	if httpRespBody == nil {
		diagnostics.AddError(AicAPIError, errorSummary+"\n"+err.Error())
	} else {
		tflog.Debug(ctx, "Error HTTP response body: "+string(httpRespBody))
		var aicError AicErrorResponse
		internalError := json.Unmarshal(httpRespBody, &aicError)
		if internalError == nil {
			var errorDetail strings.Builder
			errorDetail.WriteString("Error summary: ")
			errorDetail.WriteString(errorSummary)
			errorDetail.WriteString("\nMessage: ")
			errorDetail.WriteString(aicError.Message)
			errorDetail.WriteString("\nCode: ")
			errorDetail.WriteString(strconv.Itoa(aicError.Code))
			diagnostics.AddError(AicAPIError, errorDetail.String())
		} else {
			diagnostics.AddError(AicAPIError, errorSummary+"\n"+err.Error()+" - Detail:\n"+string(httpRespBody))
		}
	}
}
