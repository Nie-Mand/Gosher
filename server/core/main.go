package core

import (
	"context"

	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

	server := grpc.NewServer()

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
* @function: GetUser
* @description: Get the user from the context
* @params: ctx: context.Context
* @return: string
 */
func GetUser(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	if len(md["who"]) == 0 {
		return ""
	}

	return md["who"][0]
}
