package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of Gosher Client CLI",
  Long:  `All software has versions. This is Gosher's Client CLI`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Gosher Client CLI: 0.1")
  },
}