package app

import (
	"fmt"
	"strings"
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
	fmt.Println("-------------------- BUILD INFO --------------------")
	fmt.Printf(strings.TrimPrefix(buildInfoTemplate, "\n"),
		Version, Commit, Branch, BuildTime, Builder, GoVersion, LastCommitTime, LastCommitAuthor)
	fmt.Println("----------------------------------------------------")
}
