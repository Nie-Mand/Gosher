package core

import (
	"fmt"

	"Nie-Mand/Gosher/cli/schemas"

	"google.golang.org/grpc"
)

func HandleCLI(client *grpc.ClientConn, args []string) {

	api := schemas.NewGosherClient(client)

	if len(args) == 0 {
		fmt.Println("Hi, No command provided")
		return
	}
	command := args[0]
	switch command {
	case "init":
		Init()

	case "sayHi":
		SayHi(&api)

	case "ping":
		PingForFile(&api, args[1])

	case "listen":
		if len(args) > 1 {
			if args[1] == "pings" {
				ListenForFilePings(&api)
			} else if args[1] == "files" {
				SeedFile(&api)
			}
		} else {
			fmt.Println("No event provided")
		}

	case "request":
		if len(args) > 2 {
			RequestFile(&api, args[1], args[2])
		} else {
			fmt.Println("No file name or user provided")
		}

	case "receiveHi":
		ReceiveHi(&api)
	default:
		fmt.Println("unknown command")
	}

}
