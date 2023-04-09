package core

import (
	"log"
	"net"

	"Nie-Mand/Gosher/server/schemas"
	"Nie-Mand/Gosher/server/services"
	"Nie-Mand/Gosher/server/utils"

	"google.golang.org/grpc"
)

/*
* @function: CreateServer
* @description: Create a server and return the server and a function to listen to the server
* @params: port: string
* @return: *grpc.Server, func()
 */
func CreateServer(uri string) (*grpc.Server, func()) {
	listener, _error := net.Listen("tcp", uri)

	if _error != nil {
		log.Fatalf("Failed to listen: %v", _error)
	}

	opts := []grpc.ServerOption{}

	// opts = append(opts, grpc.Creds(GetTLSCredentials()))

	server := grpc.NewServer(opts...)

	log.Printf("Server listening at %v", listener.Addr())

	listen := func() {
		_error = server.Serve(listener)
		if _error != nil {
			log.Fatalf("failed to serve: %v", _error)
		}
	}

	return server, listen
}

/*
* @function: StartServer
* @description: Register the gRPC API and Start the Server
* @params: None
* @return: None
 */
func StartServer() {
	port := utils.GetEnv("SERVER_PORT", "50051")

	server, listen := CreateServer(port)

	schemas.RegisterGosherServer(server, &services.ServerStruct{})

	listen()
}

