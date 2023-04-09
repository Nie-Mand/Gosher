package core

import (
	"Nie-Mand/Gosher/cli/schemas"
	"fmt"
	"io"
	"os"
)

func PingForFile(
	api *schemas.GosherClient,
	description string,
) {
	fmt.Println("Pinging for file with description: " + description + "..")
	ctx := GetContext()

	receiver, _error := (*api).PingForFile(ctx, &schemas.PingForFileRequest{
		Description: description,
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

		if _error != nil {
			fmt.Println("Error receiving file")
			panic(_error)
		}
		fmt.Println(response.Who + " has the file: " + response.FileName)
	}

	fmt.Println("Done")
}

func ListenForFilePings(
	api *schemas.GosherClient,
) {
	fmt.Println("Listening for file pings..")
	ctx := GetContext()

	stream, _error := (*api).ListenForFilePings(ctx)
	
	if _error != nil {
		fmt.Printf("Error: %v\n", _error)
		} else {
			for {
				response, _error := stream.Recv()
				if _error != nil {
					fmt.Printf("Error: %v\n", _error)
					os.Exit(1)
				}
			fmt.Println("Somebody is looking for a file with description: " + response.Description)
			files := CheckFilesForMatch(response.Description)
			for _, file := range files {
				fmt.Println("Relevant File: " + file)
				stream.SendMsg(&schemas.ListenForFilePingsRequest{
					FileName:    file,
					Description: response.Description,
				})
			}
		}
	}
}
