package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/cloneable/repo-checker/internal/github"
)

func main() {
	ctx := context.Background()
	var (
		token = flag.String("token", "", "GitHub personal access token")
		owner = flag.String("owner", "", "login of repository owner")
	)
	flag.Parse()

	gh, err := github.New(*token)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := gh.OwnerRepos(ctx, *owner, 10, "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp.RepositoryOwner)
}
