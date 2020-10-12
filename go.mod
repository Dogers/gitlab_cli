module github.com/Dogers/gitlab_cli

go 1.13

replace github.com/Dogers/gitlab_cli/cmd => ./cmd

require (
	github.com/Dogers/gitlab_cli/cmd v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.0.0 // indirect
	github.com/xanzy/go-gitlab v0.38.1 // indirect
)
