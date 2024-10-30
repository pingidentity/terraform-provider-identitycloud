package acctest

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"testing"

	client "github.com/pingidentity/identitycloud-go-client/identitycloud"
	"github.com/pingidentity/terraform-provider-identitycloud/internal/auth"
)

var testOverrideUrlRegex = regexp.MustCompile(`^https://127.0.0.1:\d{5}$`)

// Verify that any required environment variables are set before the test begins
func ConfigurationPreCheck(t *testing.T) {
	errorFound := false
	if os.Getenv("PINGAIC_TF_TEST_MOCK_SERVICE") == "true" {
		if os.Getenv("PINGAIC_TF_TENANT_ENV_FQDN") != "" {
			t.Errorf("The 'PINGAIC_TF_TENANT_ENV_FQDN' environment variable must not be set when running acceptance tests with the mock service via setting `PINGAIC_TF_TEST_MOCK_SERIVCE` to true")
			errorFound = true
		} else {
			os.Setenv("PINGAIC_TF_ACCESS_TOKEN", "placeholder")
		}
	} else {
		envVars := []string{
			"PINGAIC_TF_TENANT_ENV_FQDN",
			"PINGAIC_TF_ACCESS_TOKEN",
		}
		for _, envVar := range envVars {
			if os.Getenv(envVar) == "" {
				t.Errorf("The '%s' environment variable must be set to run acceptance tests", envVar)
				errorFound = true
			}
		}
	}

	if errorFound {
		t.FailNow()
	}
}

// Client to be used directly in tests
func Client(testOverrideUrl *string) *client.APIClient {
	clientUrl := fmt.Sprintf("https://%s", os.Getenv("PINGAIC_TF_TENANT_ENV_FQDN"))
	if testOverrideUrl != nil && testOverrideUrlRegex.MatchString(*testOverrideUrl) {
		clientUrl = *testOverrideUrl
	}
	clientConfig := client.NewConfiguration()
	clientConfig.Servers = client.ServerConfigurations{
		{
			URL: clientUrl,
		},
	}
	clientConfig.HTTPClient = &http.Client{}
	if testOverrideUrl != nil {
		// This will only be used for acceptance tests that mock the service. The override URL is verified above.
		// #nosec G402
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		clientConfig.HTTPClient.Transport = tr
	}
	return client.NewAPIClient(clientConfig)
}

func AuthContext() context.Context {
	return auth.AuthContext(context.Background(), os.Getenv("PINGAIC_TF_ACCESS_TOKEN"))
}
