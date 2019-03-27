package cmd

import (
	"fmt"
	"os"

	zmqbroker "github.com/junjuew/mkt/broker"
	"github.com/spf13/cobra"
)

var (
	frontendURI string
	backendType string
	backendURI  string
	cfgFile     string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mkt",
	Short: "A broker that manages multiple TCP sockets",
	Long: `It accepts connections from multiple clients
and then forwards request to another endpoint either through TCP
connections or a message queue.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting broker to connect %s to %s (%s)...\n",
			frontendURI, backendURI, backendType)
		zmqbroker.CreateBroker(frontendURI, backendType, backendURI)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	rootCmd.Flags().StringVarP(&frontendURI, "furi", "f", "", "Frontend URI")
	rootCmd.MarkFlagRequired("furi")
	rootCmd.Flags().StringVarP(&backendType, "btype", "t", "", "Backend Type")
	rootCmd.MarkFlagRequired("btype")
	rootCmd.Flags().StringVarP(&backendURI, "buri", "b", "", "Backend URI")
	rootCmd.MarkFlagRequired("buri")
}
