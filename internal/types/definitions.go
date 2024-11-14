package types

import client "github.com/pingidentity/identitycloud-go-client/identitycloud"

// Configuration passed to resources
type ResourceConfiguration struct {
	ApiClient            *client.APIClient
	AccessToken          *string
	ServiceAccountConfig *client.ServiceAccountTokenSource
}
