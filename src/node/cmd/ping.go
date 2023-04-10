package cmd

import (
	"Nie-Mand/Gosher/cli/schemas"
	"Nie-Mand/Gosher/cli/utils"
	"io"
	"log"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(pingCmd)
}

var pingCmd = &cobra.Command{
  Use:   "ping [description]",
  Short: "Ping for a file",
  Long:  `Search for a file with a given description`,
  Args: cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
	client := utils.GetClient()
	api := schemas.NewGosherClient(client)
	ctx := utils.GetContext()
	description := args[0]
	log.Println("Pinging for file with description: " + description + "..")

	receiver, _error := api.PingForFile(ctx, &schemas.PingForFileRequest{
		Description: description,
	})

	if _error != nil {
		log.Printf("Error: %v\n", _error)
		panic(_error)
	}

		for {
		response, _error := receiver.Recv()
		if _error == io.EOF {
			break
		}

		if _error != nil {
			log.Println("Error receiving file")
			panic(_error)
		}
		log.Println(response.Who + " has the file: " + response.FileName)
	}

	log.Println("Done")
  },
}