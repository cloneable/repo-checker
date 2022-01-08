// Package github provides a client whose methods query the GraphQL API.
package github

import (
	"errors"
	"net/http"
	"time"

	"github.com/Khan/genqlient/graphql"
)

const (
	// githubGraphQLEndpoint is the full URL of GitHub's GraphQL API endpoint.
	githubGraphQLEndpoint = "https://api.github.com/graphql"

	// defaultHTTPRequestTimeout defines a generous default timeout.
	defaultHTTPRequestTimeout = 5 * time.Second
)

// ErrNoTokenSpecified is returned by New when no token was specified.
var ErrNoTokenSpecified = errors.New("no token specified")

// Client is used to query GitHub using pre-defined queries.
type Client struct {
	gqlClient graphql.Client
}

// New returns a new Client using the token to authenticate.
func New(personalAccessToken string) (*Client, error) {
	if personalAccessToken == "" {
		return nil, ErrNoTokenSpecified
	}
	return &Client{
		gqlClient: graphql.NewClient(githubGraphQLEndpoint, &http.Client{
			Timeout: defaultHTTPRequestTimeout,
			Transport: &authTransport{
				bearerToken: personalAccessToken,
				wrapped:     http.DefaultTransport,
			},
		}),
	}, nil
}

type authTransport struct {
	wrapped     http.RoundTripper
	bearerToken string
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "bearer "+t.bearerToken)
	return t.wrapped.RoundTrip(req)
}
