package main

import (
	"fmt"
	"github.com/google/go-github/github"
)

func outputPullRequest(pr []*github.PullRequest, repo *repositoryData) error {
	for _, p := range pr {
		num := *p.Number
		title := *p.Title
		line := fmt.Sprintf("- [#%d](https://github.com/%s/%s/pull/%d) %s", num, repo.owner, repo.repository, num, title)

		fmt.Println(line)
	}
	return nil
}
