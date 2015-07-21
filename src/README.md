# Pull-Request-Changelog
the pull-request-changelog create changelog from github pull-request data.

# Feature
* get pull-request number from commit message. 
* collect data from pull-request on github.com.
* output changelog from pull-request data.
* ~~output template support.~~ (not yet)

# Usage
pull-request-changelog -start v1.0.0 -end v2.0.0 -t GITHUB_API_TOKEN


# Output
The default output like this.
~~This software support template, so you can change output.~~(not yet)

```
- [#50](https://github.com/.Repo.owner/.Repo.repository/pull/50) change :due to :start
- [#49](https://github.com/.Repo.owner/.Repo.repository/pull/49) ls task default change
- [#48](https://github.com/.Repo.owner/.Repo.repository/pull/48) ls command support completed/uncompleted task query
- [#47](https://github.com/.Repo.owner/.Repo.repository/pull/47) add subtask test
```


# Dependency 
This software depend three Rules.

## git remote
This software use `git remote -v`.
And expect output format like `origin  git@github.com:USERNAME/REPOSITORY_NAME.git (fetch/pull)`.

This means, it works on github managed repository only.
So, not work other repository which like gitlab, bitbuket.

## git log
This software use `git log --pretty=oneline COMMIT1...COMMIT2`. 
And expect the command output format like `COMMIT_HASH COMMIT_MESSAGE`.
If your git don't return this format, this isn't work well.

## merge commit's commit message
This software expect specific format in merge commit which pull request merge commit.
The commit message should be start with `Merge pull request #NUMBER .....`.

When accept merge request, Github create merge commit with message which follow the format (2015/06/01).
If Github change merge commit message format, this will come not to do.


# How to build
cd src
gom install
gom build