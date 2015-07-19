package main

import (
	"os"
	"text/template"

	"github.com/google/go-github/github"
)

const defaultTemplate = `{{range $p := .Pr}}- [#{{$p.Number}}](https://github.com/.Repo.owner/.Repo.repository/pull/{{$p.Number}}) {{$p.Title}}
{{end}}`

type templateData struct {
	Pr   []*github.PullRequest
	Repo *repositoryData
}

func outputPullRequest(pr []*github.PullRequest, repo *repositoryData) error {
	data := &templateData{
		Pr:   pr,
		Repo: repo,
	}

	tp := template.Must(template.New("changelogTemplate").Parse(defaultTemplate))
	err := tp.Execute(os.Stdout, data)
	return err
}
