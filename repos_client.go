// Copyright (c) 2024 RFull Development
// This source code is managed under the MIT license. See LICENSE in the project root.
package main

import (
	"context"

	"github.com/octokit/go-sdk/pkg"
	"github.com/octokit/go-sdk/pkg/github/repos"
)

type ReposClient struct {
	client *repos.ItemRepoItemRequestBuilder
}

func NewReposClient(owner string, repo string) *ReposClient {
	client, e := pkg.NewApiClient(
		pkg.WithBaseUrl("https://api.github.com"),
	)
	if e != nil {
		return nil
	}
	repos := client.Repos().ByOwnerId(owner).ByRepoId(repo)
	return &ReposClient{
		client: repos,
	}
}

func (instance *ReposClient) License() (string, error) {
	license, e := instance.client.License().Get(context.Background(), nil)
	if e != nil {
		return "", e
	}
	value := *license.GetLicense().GetName()
	return value, nil
}
