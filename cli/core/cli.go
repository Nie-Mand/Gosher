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

	case "receiveHi":
		ReceiveHi(&api)
	default:
		fmt.Println("unknown command")
	}

}
