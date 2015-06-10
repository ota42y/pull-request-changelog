package main

import (
	"flag"
	"fmt"

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

func getPullRequestData(numbers []int, data *repositoryData, token string) []*github.PullRequest {
	ts := &tokenSource{
		&oauth2.Token{AccessToken: token},
	}

	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	var pls []*github.PullRequest
	for _, number := range numbers {
		repos, _, err := client.PullRequests.Get(data.owner, data.repository, number)
		if err != nil {
			fmt.Println(err)
		} else {
			pls = append(pls, repos)
		}
	}

	return pls
}

func main() {
	var remote string
	flag.StringVar(&remote, "remote", "origin", "github remote repository name (deafult: origin)")
	flag.StringVar(&remote, "r", "origin", "github remote repository name (deafult: origin)")

	var token string
	flag.StringVar(&token, "token", "", "github api token(if not set, read from pit)")
	flag.StringVar(&token, "t", "", "github api token(if not set, read from pit)")

	var start string
	flag.StringVar(&start, "start", "v0.0.1", "oldest commit setting (default: v0.0.1)")

	var end string
	flag.StringVar(&end, "end", "origin/master", "newest commit setting (default: origin/master)")

	flag.Parse()

	nums, _ := getAllPullRequestNumbers(start, end)
	data, _ := getRepositoryData(remote)
	pr := getPullRequestData(nums, data, token)

	for _, p := range pr {
		num := *p.Number
		title := *p.Title
		line := fmt.Sprintf("- [#%d](https://github.com/%s/%s/pull/%d) %s", num, data.owner, data.repository, num, title)

		fmt.Println(line)
	}
}
