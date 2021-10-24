//go:generate go run github.com/Khan/genqlient

// Package github provides a client whose methods query the GraphQL API.
package github

import (
	"context"
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

func (gh *Client) OwnerRepos(ctx context.Context, login string, repoCount int, repoCursor string) (*OwnerReposResponse, error) {
	_ = `# @genqlient
query OwnerRepos(
	$login: String!,
	$repoCount: Int!,
	# @genqlient(omitempty: true)
	$repoCursor: String) {
  repositoryOwner(login: $login) {
    repositories(first: $repoCount, after: $repoCursor) {
  	pageInfo {
  	  endCursor
  	  hasNextPage
  	}
  	edges {
  	  cursor
  	  node {
  		id
  		nameWithOwner
  		description
  		branchProtectionRules(first: 10) {
  		  nodes {
  			id
  			pattern
  			allowsDeletions
  			allowsForcePushes
  			isAdminEnforced
  			requiredApprovingReviewCount
  			requiresApprovingReviews
  			requiresLinearHistory
  			requiresStatusChecks
  			restrictsPushes
  		  }
  		}
  		autoMergeAllowed
  		defaultBranchRef {
  		  name
  		}
  		isArchived
  		isLocked
  		visibility
  		deleteBranchOnMerge
  		forkingAllowed
  		forks(first: 10) {
  		  nodes {
  			id
  			owner {
  			  login
  			}
  			name
  		  }
  		}
  		hasIssuesEnabled
  		hasWikiEnabled
  		hasProjectsEnabled
  		isPrivate
  		isTemplate
  		isSecurityPolicyEnabled
  		labels(first: 20) {
  		  nodes {
  			id
  			name
  			description
  			color
  			isDefault
  			pullRequests {
  			  totalCount
  			}
  			issues {
  			  totalCount
  			}
  		  }
  		}
  	  }
  	}
    }
  }
}`
	return OwnerRepos(ctx, gh.gqlClient, login, repoCount, repoCursor)
}
