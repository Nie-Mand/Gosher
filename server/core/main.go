package core

import (
	"os"

	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
* @function: GetEnv
* @description: Get the environment variable
* @params: envVarName: string, defaultValue: string
* @return: string
 */
func GetEnv(envVarName string, defaultValue string) string {
	port := os.Getenv(envVarName)
	if port == "" {
		port = defaultValue
	}
	return ":" + port
}

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
		log.Fatalf("fail to dial: %v", _error)
	}

	return connection
}
