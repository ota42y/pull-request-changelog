package main

import (
	"io/ioutil"
	"os"
	"text/template"

	"github.com/google/go-github/github"
)

type templateData struct {
	Pr   []*github.PullRequest
	Repo *repositoryData
}

func getTemplate(templateFile string) (string, error) {
	if templateFile != "" {
		buf, err := ioutil.ReadFile(templateFile)
		if err != nil {
			return "", err
		}

		return string(buf), nil
	}

	tpl, err := Asset("template.tpl")
	if err != nil {
		return "", err
	}

	return string(tpl), nil
}

func outputPullRequest(pr []*github.PullRequest, repo *repositoryData, templateFile string) error {
	data := &templateData{
		Pr:   pr,
		Repo: repo,
	}
	templateString, err := getTemplate(templateFile)
	if err != nil {
		return err
	}

	tp := template.Must(template.New("changelogTemplate").Parse(templateString))
	err = tp.Execute(os.Stdout, data)
	return err
}
