package github

import (
	"context"
)

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
	  # @genqlient(typename: OwnedRepo)
  	  node {
  		id
  		name
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
  	  }
  	}
    }
  }
}`
	return OwnerRepos(ctx, gh.gqlClient, login, repoCount, repoCursor)
}

func (gh *Client) RepoLabels(ctx context.Context, repoOwner, repoName string, labelCount int, labelCursor string) (*RepoLabelsResponse, error) {
	_ = `# @genqlient
		query RepoLabels(
			$repoOwner: String!,
			$repoName: String!,
			$labelCount: Int!,
			# @genqlient(omitempty: true)
			$labelCursor: String) {
			repository(owner: $repoOwner, name: $repoName) {
			  labels(first: $labelCount, after: $labelCursor) {
				pageInfo {
				  endCursor
				  hasNextPage
				}
				edges {
				  cursor
				  # @genqlient(typename: RepoLabel)
				  node {
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
		  `
	return RepoLabels(ctx, gh.gqlClient, repoOwner, repoName, labelCount, labelCursor)
}
