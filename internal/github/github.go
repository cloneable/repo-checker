//go:generate go run github.com/Khan/genqlient

package github

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type Client struct {
	gqlClient graphql.Client
}

func New(client graphql.Client) *Client {
	return &Client{gqlClient: client}
}

func (gh *Client) OwnerRepos(ctx context.Context, login string) (*OwnerReposResponse, error) {
	_ = `# @genqlient
	query OwnerRepos($login: String!) {
		repositoryOwner(login: $login) {
			repositories(first:20) {
				nodes {
					name
					createdAt
					description
				}
			}
	  	}
	}`
	return OwnerRepos(ctx, gh.gqlClient, login)
}
