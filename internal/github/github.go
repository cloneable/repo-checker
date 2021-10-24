//go:generate go run github.com/Khan/genqlient

// Package github provides a client whose methods query the GraphQL API.
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
