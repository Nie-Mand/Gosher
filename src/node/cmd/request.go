package cmd

import (
	"Nie-Mand/Gosher/cli/schemas"
	"Nie-Mand/Gosher/cli/utils"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(requestCmd)
}

var requestCmd = &cobra.Command{
	Use:   "request [file] [user]",
	Short: "Request a file from a user",
	Long:  `Request a file from a user`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		user := args[1]
		fmt.Println("Requesting the file: " + file + " from " + user + "..")
		client := utils.GetClient()
		api := schemas.NewGosherClient(client)
		ctx := utils.GetContext()

		receiver, _error := api.RequestFile(ctx, &schemas.RequestFileRequest{
			Who:      user,
			FileName: file,
		})

		if _error != nil {
			fmt.Printf("Error: %v\n", _error)
			panic(_error)
		}

		for {
			response, _error := receiver.Recv()
			if _error == io.EOF {
				break
			}

			if response.File != nil {
				utils.SaveFile("data/"+file, response.File)
				break
			} else {
				fmt.Println(response.Msg)
				break
			}
		}

		fmt.Println("Done")
	},
}
