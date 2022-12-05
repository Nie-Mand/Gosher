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
	response, _error := (*api).SayHi(ctx, &schemas.Request{
		Msg:         "BRUVVV",
		Destination: "niemand1",
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

	if err := os.Mkdir("shared", os.ModePerm); err != nil {
		fmt.Println("Error creating shared directory")
		panic(err)

	} else {
		fmt.Println("Created shared directory")
	}

	// handle TLS Stuffs
}
