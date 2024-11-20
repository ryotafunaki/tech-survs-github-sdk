// Copyright (c) 2024 RFull Development
// This source code is managed under the MIT license. See LICENSE in the project root.
package main

import (
	"fmt"
	"testing"
)

func TestRepositoryInfo(t *testing.T) {
	client := NewReposClient("ryotafunaki", "tech-survs-github-sdk")
	if client == nil {
		return
	}
	repoInfo, e := client.License()
	if e != nil {
		t.Fatalf("failed to get license: %v", e)
	}
	fmt.Println(repoInfo)
}
