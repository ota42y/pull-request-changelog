package main

import (
	"strconv"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"regexp"
)

type repositoryData struct {
	owner      string
	repository string
}

func parseRepositoryData(input io.Reader, remoteName string) *repositoryData {
	var repositoryGitAtExp, _ = regexp.Compile(fmt.Sprintf("%s\\s+git@github.com:(.+)/(.+).git.*", remoteName))
	var repositoryGitSchemeExp, _ = regexp.Compile(fmt.Sprintf("%s\\s+git://github.com/(.+)/(.+).git.*", remoteName))
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		b := []byte(line)

		match := repositoryGitAtExp.FindSubmatch(b)
		if len(match) == 3 {
			return &repositoryData{
				owner:      string(match[1]),
				repository: string(match[2]),
			}
		}

		match = repositoryGitSchemeExp.FindSubmatch(b)
		if len(match) == 3 {
			return &repositoryData{
				owner:      string(match[1]),
				repository: string(match[2]),
			}
		}
	}
	return nil
}

func getRepositoryData(remoteName string) (*repositoryData, error) {
	out, err := exec.Command("git", "remote", "-v").Output()
	if err != nil {
		return nil, err
	}

	return parseRepositoryData(bytes.NewBuffer(out), remoteName), nil
}

// getAllPullRequestNumbers return all pull request id
// This method support github pull request style only
func getAllPullRequestNumbers(from string, to string) ([]int, error) {
	out, err := exec.Command("git", "log", "--pretty=oneline", from+"..."+to).Output()
	if err != nil {
		return make([]int, 0), err
	}
	return parsePullRequestNumbers(bytes.NewBuffer(out))
}

func parsePullRequestNumbers(input io.Reader) ([]int, error) {
	var nums []int

	mergeRequestRegexp, err := regexp.Compile("^[a-f0-9]+ Merge pull request #([0-9]+) .*")
	if err != nil {
		return nums, err
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		match := mergeRequestRegexp.FindSubmatch([]byte(line))
		if 1 < len(match) {
			n, err := strconv.Atoi(string(match[1]))
			if err != nil {
				panic(err)
			}
			nums = append(nums, n)
		}
	}
	return nums, nil

}
