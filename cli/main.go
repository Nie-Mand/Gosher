package main

import (
	"Nie-Mand/Gosher/cli/core"
	"os"
)

func main() {
	core.LoadDotEnv()

	host := core.GetEnv("SERVER_HOST", "localhost")
	port := core.GetEnv("SERVER_PORT", "50051")
	uri := host + ":" + port

	client := core.CreateClient(uri)

	defer client.Close()

	core.HandleCLI(client, os.Args[1:])
}
