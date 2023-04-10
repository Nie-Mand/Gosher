package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
  Use:   "init",
  Short: "Initialize Gosher Client CLI data directory",
  Long:  `Gosher Client CLI needs a data directory to store its data`,
  Run: func(cmd *cobra.Command, args []string) {
      if err := os.Mkdir("data", os.ModePerm); err != nil {
        fmt.Println("Error creating data directory")
        panic(err)
    
      } else {
        fmt.Println("Created data directory")
      }
  },
}