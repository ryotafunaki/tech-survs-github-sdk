// Copyright (c) 2024 RFull Development
// This source code is managed under the MIT license. See LICENSE in the project root.
package main

import (
	"flag"
	"fmt"
)

func main() {
	var owner string
	var repos string
	const (
		ownerUsage = "owner of the repository"
		reposUsage = "name of the repository"
	)
	flag.StringVar(&owner, "owner", "", ownerUsage)
	flag.StringVar(&owner, "o", "", ownerUsage+" (shorthand)")
	flag.StringVar(&repos, "repos", "", reposUsage)
	flag.StringVar(&repos, "r", "", reposUsage+" (shorthand)")
	flag.Parse()
	if owner == "" || repos == "" {
		flag.Usage()
		return
	}

	client := NewReposClient(owner, repos)
	if client == nil {
		m := fmt.Errorf("failed to create client")
		fmt.Println(m)
		return
	}
	license, e := client.License()
	if e != nil {
		m := fmt.Errorf("failed to get license: %v", e).Error()
		fmt.Println(m)
		return
	}
	fmt.Println(license)
}
