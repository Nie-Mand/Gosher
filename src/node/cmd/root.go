package cmd

import (
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gosher client CLI",
	Short: "Gosher is a file sharing protocol built on top of gRPC",
	Long:  `Gosher is a file sharing protocol built on top of gRPC`,
	PreRun: func(cmd *cobra.Command, args []string) {
		ctx := context.WithValue(cmd.Context(), "stuff", "WORLD")
		cmd.SetContext(ctx)
	},
	Run:   func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
