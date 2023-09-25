package main

import "github.com/bogdzn/razin/server/cmd"

func main() {

	cmd := cmd.LoadServerCLI()
	cmd.ExecuteC()
}
