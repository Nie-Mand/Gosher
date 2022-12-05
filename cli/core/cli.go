package core

import (
	"fmt"

	"google.golang.org/grpc"
)

func HandleCLI(client *grpc.ClientConn, args []string) {

	if len(args) == 0 {
		fmt.Println("Hi, No command provided")
		return
	}
	command := args[0]
	switch command {
	case "init":
		Init()

	case "sayHi":
		SayHi(client)
	default:
		fmt.Println("unknown command")
	}

}
