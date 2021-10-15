package main

import (
	"server/core/cmd"
)

var buildTime, version, commit string

func main() {
	cmd.Execute(buildTime, version, commit)
}
