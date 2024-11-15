package acctest

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	client "github.com/pingidentity/identitycloud-go-client/identitycloud"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/auth"
)

// Verify that any required environment variables are set before the test begins
func ConfigurationPreCheck(t *testing.T) {
	errorFound := false
	envVars := []string{
		"PINGAIC_TF_TENANT_ENV_FQDN",
		"PINGAIC_TF_SERVICE_ACCOUNT_ID",
		"PINGAIC_TF_SERVICE_ACCOUNT_PRIVATE_KEY",
	}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			t.Errorf("The '%s' environment variable must be set to run acceptance tests", envVar)
			errorFound = true
		}
	}

	if os.Getenv("PINGAIC_TF_ACCESS_TOKEN") != "" {
		t.Errorf("The 'PINGAIC_TF_ACCESS_TOKEN' environment variable must not be set to run acceptance tests")
		errorFound = true
	}

	if errorFound {
		t.FailNow()
	}
}

// Client to be used directly in tests
func Client() *client.APIClient {
	clientUrl := fmt.Sprintf("https://%s", os.Getenv("PINGAIC_TF_TENANT_ENV_FQDN"))
	clientConfig := client.NewConfiguration()
	clientConfig.Servers = client.ServerConfigurations{
		{
			URL: clientUrl,
		},
	}
	clientConfig.HTTPClient = &http.Client{}
	return client.NewAPIClient(clientConfig)
}

func AuthContext() context.Context {
	tokenSource := client.ServiceAccountTokenSource{
		TenantFqdn:               os.Getenv("PINGAIC_TF_TENANT_ENV_FQDN"),
		ServiceAccountId:         os.Getenv("PINGAIC_TF_SERVICE_ACCOUNT_ID"),
		ServiceAccountPrivateKey: os.Getenv("PINGAIC_TF_SERVICE_ACCOUNT_PRIVATE_KEY"),
		Scopes: []string{
			"fr:idc:certificate:*",
			"fr:idc:content-security-policy:*",
			"fr:idc:cookie-domain:*",
			"fr:idc:custom-domain:*",
			"fr:idc:esv:* fr:idc:promotion:*",
			"fr:idc:sso-cookie:*",
		},
	}
	return auth.AuthContext(context.Background(), nil, &tokenSource)
}
