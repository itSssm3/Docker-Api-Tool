package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	address       string
	clientversion string
	proxyaddr     string
	imagename     string
	command       string
	containerid   string
)

var rootCmd = &cobra.Command{
	Use: "Docker-Api-Tool",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&proxyaddr, "proxyaddr", "", "Set proxy: socks5://127.0.0.1:6666 (optional)")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SilenceErrors = true
}
