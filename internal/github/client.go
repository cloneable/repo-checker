// Package github provides a client whose methods query the GraphQL API.
package github

import (
	"errors"
	"net/http"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// A generous default timeout.
const defaultHTTPRequestTimeout = 5 * time.Second

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
		gqlClient: graphql.NewClient("https://api.github.com/graphql", &httpClient{
			client: http.Client{
				Timeout: defaultHTTPRequestTimeout,
			},
			bearerToken: personalAccessToken,
		}),
	}, nil
}
