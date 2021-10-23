package github

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Khan/genqlient/graphql"
)

// OwnerReposRepositoryOwner includes the requested fields of the GraphQL interface RepositoryOwner.
//
// OwnerReposRepositoryOwner is implemented by the following types:
// OwnerReposRepositoryOwnerOrganization
// OwnerReposRepositoryOwnerUser
// The GraphQL type's documentation follows.
//
// Represents an owner of a Repository.
type OwnerReposRepositoryOwner interface {
	implementsGraphQLInterfaceOwnerReposRepositoryOwner()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetRepositories returns the interface-field "repositories" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// A list of repositories that the user owns.
	GetRepositories() OwnerReposRepositoryOwnerRepositoriesRepositoryConnection
}

func (v *OwnerReposRepositoryOwnerOrganization) implementsGraphQLInterfaceOwnerReposRepositoryOwner() {
}

// GetTypename is a part of, and documented with, the interface OwnerReposRepositoryOwner.
func (v *OwnerReposRepositoryOwnerOrganization) GetTypename() string { return v.Typename }

// GetRepositories is a part of, and documented with, the interface OwnerReposRepositoryOwner.
func (v *OwnerReposRepositoryOwnerOrganization) GetRepositories() OwnerReposRepositoryOwnerRepositoriesRepositoryConnection {
	return v.Repositories
}

func (v *OwnerReposRepositoryOwnerUser) implementsGraphQLInterfaceOwnerReposRepositoryOwner() {}

// GetTypename is a part of, and documented with, the interface OwnerReposRepositoryOwner.
func (v *OwnerReposRepositoryOwnerUser) GetTypename() string { return v.Typename }

// GetRepositories is a part of, and documented with, the interface OwnerReposRepositoryOwner.
func (v *OwnerReposRepositoryOwnerUser) GetRepositories() OwnerReposRepositoryOwnerRepositoriesRepositoryConnection {
	return v.Repositories
}

func __unmarshalOwnerReposRepositoryOwner(b []byte, v *OwnerReposRepositoryOwner) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Organization":
		*v = new(OwnerReposRepositoryOwnerOrganization)
		return json.Unmarshal(b, *v)
	case "User":
		*v = new(OwnerReposRepositoryOwnerUser)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"Response was missing RepositoryOwner.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for OwnerReposRepositoryOwner: "%v"`, tn.TypeName)
	}
}

func __marshalOwnerReposRepositoryOwner(v *OwnerReposRepositoryOwner) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *OwnerReposRepositoryOwnerOrganization:
		typename = "Organization"

		result := struct {
			TypeName string `json:"__typename"`
			*OwnerReposRepositoryOwnerOrganization
		}{typename, v}
		return json.Marshal(result)
	case *OwnerReposRepositoryOwnerUser:
		typename = "User"

		result := struct {
			TypeName string `json:"__typename"`
			*OwnerReposRepositoryOwnerUser
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`Unexpected concrete type for OwnerReposRepositoryOwner: "%T"`, v)
	}
}

// OwnerReposRepositoryOwnerOrganization includes the requested fields of the GraphQL type Organization.
// The GraphQL type's documentation follows.
//
// An account on GitHub, with one or more owners, that has repositories, members and teams.
type OwnerReposRepositoryOwnerOrganization struct {
	Typename string `json:"__typename"`
	// A list of repositories that the user owns.
	Repositories OwnerReposRepositoryOwnerRepositoriesRepositoryConnection `json:"repositories"`
}

// OwnerReposRepositoryOwnerRepositoriesRepositoryConnection includes the requested fields of the GraphQL type RepositoryConnection.
// The GraphQL type's documentation follows.
//
// A list of repositories owned by the subject.
type OwnerReposRepositoryOwnerRepositoriesRepositoryConnection struct {
	// A list of nodes.
	Nodes []OwnerReposRepositoryOwnerRepositoriesRepositoryConnectionNodesRepository `json:"nodes"`
}

// OwnerReposRepositoryOwnerRepositoriesRepositoryConnectionNodesRepository includes the requested fields of the GraphQL type Repository.
// The GraphQL type's documentation follows.
//
// A repository contains the content for a project.
type OwnerReposRepositoryOwnerRepositoriesRepositoryConnectionNodesRepository struct {
	// The name of the repository.
	Name string `json:"name"`
	// Identifies the date and time when the object was created.
	CreatedAt time.Time `json:"createdAt"`
	// The description of the repository.
	Description string `json:"description"`
}

// OwnerReposRepositoryOwnerUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A user is an individual's account on GitHub that owns repositories and can make new content.
type OwnerReposRepositoryOwnerUser struct {
	Typename string `json:"__typename"`
	// A list of repositories that the user owns.
	Repositories OwnerReposRepositoryOwnerRepositoriesRepositoryConnection `json:"repositories"`
}

// OwnerReposResponse is returned by OwnerRepos on success.
type OwnerReposResponse struct {
	// Lookup a repository owner (ie. either a User or an Organization) by login.
	RepositoryOwner OwnerReposRepositoryOwner `json:"-"`
}

func (v *OwnerReposResponse) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*OwnerReposResponse
		RepositoryOwner json.RawMessage `json:"repositoryOwner"`
		graphql.NoUnmarshalJSON
	}
	firstPass.OwnerReposResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.RepositoryOwner
		src := firstPass.RepositoryOwner
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalOwnerReposRepositoryOwner(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal OwnerReposResponse.RepositoryOwner: %w", err)
			}
		}
	}
	return nil
}

type __premarshalOwnerReposResponse struct {
	RepositoryOwner json.RawMessage `json:"repositoryOwner"`
}

func (v *OwnerReposResponse) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *OwnerReposResponse) __premarshalJSON() (*__premarshalOwnerReposResponse, error) {
	var retval __premarshalOwnerReposResponse

	{

		dst := &retval.RepositoryOwner
		src := v.RepositoryOwner
		var err error
		*dst, err = __marshalOwnerReposRepositoryOwner(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"Unable to marshal OwnerReposResponse.RepositoryOwner: %w", err)
		}
	}
	return &retval, nil
}

// __OwnerReposInput is used internally by genqlient
type __OwnerReposInput struct {
	Login string `json:"login"`
}

func OwnerRepos(
	ctx context.Context,
	client graphql.Client,
	login string,
) (*OwnerReposResponse, error) {
	__input := __OwnerReposInput{
		Login: login,
	}
	var err error

	var retval OwnerReposResponse
	err = client.MakeRequest(
		ctx,
		"OwnerRepos",
		`
query OwnerRepos ($login: String!) {
	repositoryOwner(login: $login) {
		__typename
		repositories(first: 20) {
			nodes {
				name
				createdAt
				description
			}
		}
	}
}
`,
		&retval,
		&__input,
	)
	return &retval, err
}
