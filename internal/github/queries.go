package github

import (
	"context"

	"github.com/cloneable/repo-checker/internal/api"
)

func (gh *Client) OwnerRepos(ctx context.Context, login string, repoCount int, repoCursor string) (*api.OwnerReposResponse, error) {
	return api.OwnerRepos(ctx, gh.gqlClient, login, repoCount, repoCursor)
}

func (gh *Client) RepoLabels(ctx context.Context, repoOwner, repoName string, labelCount int, labelCursor string) (*api.RepoLabelsResponse, error) {
	return api.RepoLabels(ctx, gh.gqlClient, repoOwner, repoName, labelCount, labelCursor)
}
