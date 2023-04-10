package cmd

import (
	"Nie-Mand/Gosher/cli/schemas"
	"Nie-Mand/Gosher/cli/utils"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(listenCmd)
}

var listenCmd = &cobra.Command{
  Use:   "listen [event]",
  Short: "Listen for requests",
  Long:  `Listen for ping or file requests`,
  Args: cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
	client := utils.GetClient()
	api := schemas.NewGosherClient(client)
	ctx := utils.GetContext()
	event := args[0]

	if event == "pings" {

		fmt.Println("Listening for file pings..")
		stream, _error := api.ListenForFilePings(ctx)
		if _error != nil {
			fmt.Printf("Error: %v\n", _error)
			} else {
				for {
					response, _error := stream.Recv()
					if _error != nil {
						log.Printf("Error: %v\n", _error)
						os.Exit(1)
					}
				fmt.Println("Somebody is looking for a file with description: " + response.Description)
				files := utils.CheckFilesForMatch(response.Description)
				for _, file := range files {
					log.Println("Relevant File: " + file)
					stream.SendMsg(&schemas.ListenForFilePingsRequest{
						FileName:    file,
						Description: response.Description,
					})
				}
			}
		}

	} else if event == "files" {
		fmt.Println("Listening for file requests..")
		ctx := utils.GetContext()

		stream, _error := api.SeedFile(ctx)
	if _error != nil {
		log.Printf("Error: %v\n", _error)
	} else {
		for {
			response, _error := stream.Recv()
			if _error != nil {
				log.Printf("Error: %v\n", _error)
				os.Exit(1)
			}
			log.Println("i got", response)
			filename := response.FileName

			fmt.Println("someone is requesting the file: " + filename + ", approve? [y/n]")

			var yN string
			fmt.Scanf("%s", &yN)

			if yN == "y" {
				file, err := utils.OpenFile("data/" + filename)

				if err == nil {
					stream.Send(&schemas.SeedFileRequest{
						File: file,
						Msg:  "",
					})
				}
				fmt.Println("Done")
			} else {

				stream.Send(&schemas.SeedFileRequest{
					File: nil,
					Msg:  "The owner did not approve your request",
				})

			}

			files := utils.CheckFilesForMatch(response.FileName)
			for _, file := range files {
				fmt.Println("Relevant File: " + file)
				stream.SendMsg(&schemas.ListenForFilePingsRequest{
					FileName:    file,
					Description: response.FileName,
				})
			}
		}
	}
	} else {
		log.Println("Invalid event")
	}
  },
}