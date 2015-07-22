{{range $p := .Pr}}- [#{{$p.Number}}](https://github.com/.Repo.owner/.Repo.repository/pull/{{$p.Number}}) {{$p.Title}}
{{end}}