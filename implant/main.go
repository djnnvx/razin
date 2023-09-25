package main

import "github.com/bogdzn/razin/implant/cmd"

func main() {

	cmd := cmd.LoadClientCLI()
	cmd.ExecuteC()
}
