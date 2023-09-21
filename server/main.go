package main

import "github.com/bogdzn/razin/server/cmd"

func main() {

	defaults := cmd.CliDefaults()
	cmd := cmd.LoadServerCLI(defaults)
	cmd.ExecuteC()
}
