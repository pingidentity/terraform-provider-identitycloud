package resource

import (
	"context"

	client "github.com/pingidentity/identitycloud-go-client/identitycloud"
)

func AuthContext(ctx context.Context, accessToken string) context.Context {
	return context.WithValue(ctx, client.ContextAccessToken, accessToken)
}
