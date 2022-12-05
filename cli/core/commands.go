package core

import (
	"Nie-Mand/Gosher/cli/schemas"
	"fmt"
	"os"
)

func SayHi(
	api *schemas.GosherClient,
) {
	fmt.Println("Sending Hi..")
	ctx := GetContext()

	filePath := "tmp/dev.png"
	file := GetFile(filePath)

	response, _error := (*api).SayHi(ctx, &schemas.Request{
		Msg:         "BRUVVV",
		Name:        "dev.png",
		Destination: "niemand1",
		File:        file,
	})

	if _error != nil {
		fmt.Printf("Error: %v\n", _error)
	} else {
		fmt.Printf("Response: %v\n", response.Msg)
	}

	fmt.Println("Done")
}

func ReceiveHi(
	api *schemas.GosherClient,
) {
	fmt.Println("Receiving Hi..")
	ctx := GetContext()

	stream, _error := (*api).ReceiveHi(ctx, &schemas.Identity{})
	if _error != nil {
		fmt.Printf("Error: %v\n", _error)
	} else {
		for {
			response, _error := stream.Recv()
			if _error != nil {
				fmt.Printf("Error: %v\n", _error)
				os.Exit(1)
			}
			fmt.Printf("Response: %v\n", response.Msg)
			SaveFile("downloads/"+response.Name, response.File)
		}
	}
}

func Init() {
	if err := os.Mkdir("cache", os.ModePerm); err != nil {
		fmt.Println("Error creating cache directory")
		panic(err)
	} else {
		fmt.Println("Created cache directory")
	}

	if err := os.Mkdir("downloads", os.ModePerm); err != nil {
		fmt.Println("Error creating downloads directory")
		panic(err)

	} else {
		fmt.Println("Created downloads directory")
	}

	// handle TLS Stuffs
}
