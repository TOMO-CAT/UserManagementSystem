package app

import (
	"fmt"
)

var (
	Version          string
	Commit           string
	Branch           string
	BuildTime        string
	Builder          string
	GoVersion        string
	LastCommitTime   string
	LastCommitAuthor string
)

var buildInfoTemplate = `
Version:           %s
Commit:            %s
Branch:            %s
BuildTime:         %s
Builder:           %s
GoVersion:         %s
LastCommitTime:    %s
LastCommitAuthor:  %s
`

func printBuildInfo() {
	fmt.Printf(buildInfoTemplate, Version, Commit, Branch, BuildTime, Builder, GoVersion, LastCommitTime, LastCommitAuthor)
}
