module github.com/Dogers/gitlab_cli

go 1.13

replace github.com/Dogers/gitlab_cli/cmd => ./cmd

require (
	github.com/spf13/cobra v0.0.7
	github.com/xanzy/go-gitlab v0.33.0
)
