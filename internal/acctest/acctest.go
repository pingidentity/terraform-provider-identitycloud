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
	envVars := []string{
		"PINGAIC_TF_TENANT_ENV_FQDN",
		"PINGAIC_TF_ACCESS_TOKEN",
	}

	errorFound := false
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			t.Errorf("The '%s' environment variable must be set to run acceptance tests", envVar)
			errorFound = true
		}
	}

	if errorFound {
		t.FailNow()
	}
}

// Client to be used directly in tests
func Client() *client.APIClient {
	tenantFqdn := os.Getenv("PINGAIC_TF_TENANT_ENV_FQDN")
	clientConfig := client.NewConfiguration()
	clientConfig.Servers = client.ServerConfigurations{
		{
			URL: fmt.Sprintf("https://%s", tenantFqdn),
		},
	}

	clientConfig.HTTPClient = &http.Client{}
	return client.NewAPIClient(clientConfig)
}

func AuthContext() context.Context {
	return auth.AuthContext(context.Background(), os.Getenv("PINGAIC_TF_ACCESS_TOKEN"))
}
