// Copyright © 2025 Ping Identity Corporation

package auth

import (
	"context"

	client "github.com/pingidentity/identitycloud-go-client/identitycloud"
)

func AuthContext(ctx context.Context, accessToken *string, serviceAccountConfig *client.ServiceAccountTokenSource) context.Context {
	if accessToken != nil {
		return AccessTokenAuthContext(ctx, *accessToken)
	}
	if serviceAccountConfig != nil {
		return ServiceAccountAuthContext(ctx, *serviceAccountConfig)
	}
	// This should never happen, since auth vars are verified in provider.go
	return context.Background()
}

func ServiceAccountAuthContext(ctx context.Context, serviceAccountConfig client.ServiceAccountTokenSource) context.Context {
	return context.WithValue(ctx, client.ContextOAuth2, serviceAccountConfig)
}

func AccessTokenAuthContext(ctx context.Context, accessToken string) context.Context {
	return context.WithValue(ctx, client.ContextAccessToken, accessToken)
}
