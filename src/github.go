package main

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// create struct for the token source
type tokenSource struct {
	token *oauth2.Token
}

// add Token() method to satisfy oauth2.TokenSource interface
func (t *tokenSource) Token() (*oauth2.Token, error) {
	return t.token, nil
}

// getPullRequestData return some []*github.PullRequest.
func getPullRequestData(numbers []int, data *repositoryData, accessToken string) ([]*github.PullRequest, error) {
	ts := &tokenSource{
		&oauth2.Token{AccessToken: accessToken},
	}

	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	var pls []*github.PullRequest
	for _, number := range numbers {
		repos, _, err := client.PullRequests.Get(data.owner, data.repository, number)
		if err != nil {
			return make([]*github.PullRequest, 0), err
		}
		pls = append(pls, repos)
	}

	return pls, nil
}
