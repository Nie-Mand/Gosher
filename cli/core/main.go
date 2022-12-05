package core

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

/*
* @function: CreateClient
* @description: Create a client and return the client
* @params: uri: string
* @return: *grpc.ClientConn
 */
func CreateClient(uri string) *grpc.ClientConn {
	var options []grpc.DialOption

	tls := GetEnv("SERVER_TLS", "1")

	if tls == "1" {
		// if *caFile == "" {
		// 	*caFile = data.Path("x509/ca_cert.pem")
		// }
		// creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		// if err != nil {
		// 	log.Fatalf("Failed to create TLS credentials %v", err)
		// }
		// opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	connection, _error := grpc.Dial(uri, options...)
	if _error != nil {
		fmt.Printf("Failed to connect to server: %v\n", _error)
		panic(_error)
	}

	return connection
}

/*
* @function: GetContext
* @description: Get a context
* @params: void
* @return: context.Context
 */
func GetContext() context.Context {
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(
		ctx,
		metadata.Pairs("who", GetEnv("WHO", "anonymous")),
	)
	return ctx
}
