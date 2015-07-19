package main

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetRepositoryData(t *testing.T) {
	Convey("correct", t, func() {
		Convey("git atmark", func() {
			testData := `origin  git@github.com:ota42y/plaintodo.git (fetch)
origin  git@github.com:ota42y/plaintodo.git (push)
`
			repository := parseRepositoryData(bytes.NewBufferString(testData), "origin")
			So(repository, ShouldNotBeNil)
			So(repository.owner, ShouldEqual, "ota42y")
			So(repository.repository, ShouldEqual, "plaintodo")

		})

		Convey("git scheme", func() {
			testData := `ota42y  git://github.com/ota42y/plaintodo.git (fetch)
ota42y  git://github.com/ota42y/plaintodo.git (push)
`
			repository := parseRepositoryData(bytes.NewBufferString(testData), "ota42y")
			So(repository, ShouldNotBeNil)
			So(repository.owner, ShouldEqual, "ota42y")
			So(repository.repository, ShouldEqual, "plaintodo")
		})

		Convey("multi remote", func() {
			testData := `origin  git@github.com:ota42y/plaintodo.git (fetch)
origin  git@github.com:ota42y/plaintodo.git (push)
ota42y  git://github.com/ota42y/plaintodo.git (fetch)
ota42y  git://github.com/ota42y/plaintodo.git (push)
`
			repository := parseRepositoryData(bytes.NewBufferString(testData), "ota42y")
			So(repository, ShouldNotBeNil)
			So(repository.owner, ShouldEqual, "ota42y")
			So(repository.repository, ShouldEqual, "plaintodo")
		})
	})
	Convey("incorrect", t, func() {
		Convey("not exit", func() {
			testData := `ota42y  git://github.com/ota42y/plaintodo.git (fetch)
ota42y  git://github.com/ota42y/plaintodo.git (push)
`
			repository := parseRepositoryData(bytes.NewBufferString(testData), "origin")
			So(repository, ShouldBeNil)
		})
	})
}

func TestParsePullRequestNumbers(t *testing.T) {
	Convey("correct", t, func() {
		testData := `aebddf11960f5fa4e16c3f9bf3ce3dc4674418cc Merge pull request #42 from ota42y/feature/version_up
8d0feaa0738f248599a99d5f3c05077542be5f75 v0.0.4
898e11b028f89546b44e4f27aaa0cd0a34256ca0 Merge pull request #41 from ota42y/feature/subtask_bugfix
846d44dba2e5dbb94705f5620e043448634b8060 add subtask bugfix
450ff2e7cc1e45d0f3b6ee143467666281a3ba2c Merge pull request #40 from ota42y/feature/repeat
1ab7fbfded939c002306f0404278338d6ed24911 add repeat attribute
73370fb3c4634168284c2b621e3305f1229c42ff Merge pull request #39 from ota42y/feature/task_copy
c6734cf7e6db20152144ad681209e40766388797 add task copy
c58be8e670c003c92b195e8eed0de115199aad2e Merge pull request #38 from ota42y/feature/equal_method
`
		numbers, err := parsePullRequestNumbers(bytes.NewBufferString(testData))

		So(err, ShouldBeNil)
		So(numbers, ShouldNotBeEmpty)
		So(numbers, ShouldResemble, []int{42, 41, 40, 39, 38})
	})

	Convey("invalid", t, func() {
		Convey("other pullrequest format", func() {
			testData := `aebddf11960f5fa4e16c3f9bf3ce3dc4674418cc Merge merge request #42 from ota42y/feature/version_up
8d0feaa0738f248599a99d5f3c05077542be5f75 v0.0.4
898e11b028f89546b44e4f27aaa0cd0a34256ca0 Merge merge request #41 from ota42y/feature/subtask_bugfix
846d44dba2e5dbb94705f5620e043448634b8060 add subtask bugfix
`
			numbers, err := parsePullRequestNumbers(bytes.NewBufferString(testData))
			So(err, ShouldBeNil)
			So(numbers, ShouldBeEmpty)
		})

		Convey("other input format", func() {
			testData := `Merge pull request #42 from ota42y/feature/version_up
v0.0.4
Merge pull request #41 from ota42y/feature/subtask_bugfix
add subtask bugfix
`
			numbers, err := parsePullRequestNumbers(bytes.NewBufferString(testData))
			So(err, ShouldBeNil)
			So(numbers, ShouldBeEmpty)
		})

		Convey("no input", func() {
			testData := ""
			numbers, err := parsePullRequestNumbers(bytes.NewBufferString(testData))
			So(err, ShouldBeNil)
			So(numbers, ShouldBeEmpty)
		})
	})
}
