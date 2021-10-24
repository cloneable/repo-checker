package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/cloneable/repo-checker/internal/github"
)

type httpClient struct {
	client      http.Client
	bearerToken string
}

func (c *httpClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+c.bearerToken)
	return c.client.Do(req)
}

func main() {
	ctx := context.Background()
	var (
		token = flag.String("token", "", "GitHub personal access token")
		owner = flag.String("owner", "", "login of repository owner")
	)
	flag.Parse()

	client := &httpClient{
		client: http.Client{
			Timeout: 5 * time.Second,
		},
		bearerToken: *token,
	}
	gql := graphql.NewClient("https://api.github.com/graphql", client)
	gh := github.New(gql)

	resp, err := gh.OwnerRepos(ctx, *owner, 10, "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp.RepositoryOwner)
}
