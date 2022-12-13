package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type ClientSecretAuthorizerOptions struct {
	// TODO: document these

	Environment  environments.Environment
	Api          environments.Api
	TenantId     string
	AuxTenantIds []string
	ClientId     string
	ClientSecret string
}

// NewClientSecretAuthorizer returns an authorizer which uses client secret authentication.
func NewClientSecretAuthorizer(ctx context.Context, options ClientSecretAuthorizerOptions) (Authorizer, error) {
	scope, err := environments.Scope(options.Api)
	if err != nil {
		return nil, fmt.Errorf("determining scope for %q: %+v", options.Api.Name(), err)
	}

	conf := clientCredentialsConfig{
		Environment:        options.Environment,
		TenantID:           options.TenantId,
		AuxiliaryTenantIDs: options.AuxTenantIds,
		ClientID:           options.ClientId,
		ClientSecret:       options.ClientSecret,
		Scopes: []string{
			*scope,
		},
	}

	return conf.TokenSource(ctx, clientCredentialsSecretType)
}

var _ Authorizer = &clientSecretAuthorizer{}

type clientSecretAuthorizer struct {
	conf *clientCredentialsConfig
}

func (a *clientSecretAuthorizer) Token(ctx context.Context, _ *http.Request) (*oauth2.Token, error) {
	if a.conf == nil {
		return nil, fmt.Errorf("could not request token: conf is nil")
	}

	v := url.Values{
		"client_id":     {a.conf.ClientID},
		"client_secret": {a.conf.ClientSecret},
		"grant_type":    {"client_credentials"},
		// NOTE: at this time we only support v2 (MSAL) Tokens since v1 (ADAL) is EOL.
		"scope": []string{
			strings.Join(a.conf.Scopes, " "),
		},
	}

	tokenUrl := a.conf.TokenURL
	if tokenUrl == "" {
		if a.conf.Environment.Authorization == nil {
			return nil, fmt.Errorf("no `authorization` configuration was found for this environment")
		}
		tokenUrl = tokenEndpoint(*a.conf.Environment.Authorization, a.conf.TenantID)
	}

	return clientCredentialsToken(ctx, tokenUrl, &v)
}

// AuxiliaryTokens returns additional tokens for auxiliary tenant IDs, for use in multi-tenant scenarios
func (a *clientSecretAuthorizer) AuxiliaryTokens(ctx context.Context, _ *http.Request) ([]*oauth2.Token, error) {
	if a.conf == nil {
		return nil, fmt.Errorf("could not request token: conf is nil")
	}

	tokens := make([]*oauth2.Token, 0)

	if len(a.conf.AuxiliaryTenantIDs) == 0 {
		return tokens, nil
	}

	for _, tenantId := range a.conf.AuxiliaryTenantIDs {
		v := url.Values{
			"client_id":     {a.conf.ClientID},
			"client_secret": {a.conf.ClientSecret},
			"grant_type":    {"client_credentials"},
			// NOTE: at this time we only support v2 (MSAL) Tokens since v1 (ADAL) is EOL.
			"scope": []string{
				// TODO: given the Request, could we use a dynamic scope?
				strings.Join(a.conf.Scopes, " "),
			},
		}

		tokenUrl := a.conf.TokenURL
		if tokenUrl == "" {
			if a.conf.Environment.Authorization == nil {
				return nil, fmt.Errorf("no `authorization` configuration was found for this environment")
			}
			tokenUrl = tokenEndpoint(*a.conf.Environment.Authorization, tenantId)
		}

		token, err := clientCredentialsToken(ctx, tokenUrl, &v)
		if err != nil {
			return tokens, err
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}
