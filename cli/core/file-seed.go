package core

import (
	"Nie-Mand/Gosher/cli/schemas"
	"fmt"
	"io"
	"os"
)

func RequestFile(
	api *schemas.GosherClient,
	fileName string,
	who string,
) {
	fmt.Println("Requesting the file: " + fileName + " from " + who + "..")
	ctx := GetContext()

	receiver, _error := (*api).RequestFile(ctx, &schemas.RequestFileRequest{
		Who:      who,
		FileName: fileName,
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
			SaveFile("downloads/"+fileName, response.File)
			break
		} else {
			fmt.Println(response.Msg)
			break
		}
	}

	fmt.Println("Done")
}

func SeedFile(
	api *schemas.GosherClient,
) {
	fmt.Println("Listening for file requests..")
	ctx := GetContext()

	stream, _error := (*api).SeedFile(ctx)
	if _error != nil {
		fmt.Printf("Error: %v\n", _error)
	} else {
		for {
			response, _error := stream.Recv()
			if _error != nil {
				fmt.Printf("Error: %v\n", _error)
				os.Exit(1)
			}
			fmt.Println("i got", response)
			filename := response.FileName

			fmt.Println("someone is requesting the file: " + filename + ", approve? [y/n]")

			var yN string
			fmt.Scanf("%s", &yN)

			if yN == "y" {
				file, err := OpenFile("downloads/" + filename)

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

			// files := CheckFilesForMatch(response.Description)
			// for _, file := range files {
			// 	fmt.Println("Relevant File: " + file)
			// 	stream.SendMsg(&schemas.ListenForFilePingsRequest{
			// 		FileName:    file,
			// 		Description: response.Description,
			// 	})
			// }
		}
	}
}
