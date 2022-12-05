package core

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	} else {
		fmt.Printf("Connected to server: %v\n", uri)
	}

	return connection
}
