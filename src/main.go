package main

import (
	"flag"
	"fmt"
	"os"
)

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

	if token == "" {
		fmt.Println("need github api token")
		os.Exit(1)
	}

	nums, _ := getAllPullRequestNumbers(start, end)
	repo, _ := getRepositoryData(remote)
	pr, err := getPullRequestData(nums, repo, token)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = outputPullRequest(pr, repo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
