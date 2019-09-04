package main

import "cmd"

var (
	version string
	commit  string
	date    string
)

/*

Example commands

$ certify show chain [--all] -f *.pem // Displays the complete

*/

func main() {
	cmd.Execute(version, commit, date)
}
