package utils

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

/*
* @function: GetClient
* @description: Create a client and return the client
* @params: None
* @return: *grpc.ClientConn
 */
func GetClient() (*grpc.ClientConn) {
	host := GetEnv("SERVER_HOST", "localhost")
	port := GetEnv("SERVER_PORT", "50051")
	// tls := GetEnv("SERVER_TLS", "1")

	uri := host + ":" + port

	fmt.Println("Connecting to server at: ", uri)
	var options []grpc.DialOption

	// if tls == "1" {
		// if *caFile == "" {
		// 	*caFile = data.Path("x509/ca_cert.pem")
		// }
		// creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		// if err != nil {
		// 	log.Fatalf("Failed to create TLS credentials %v", err)
		// }
		// opts = append(opts, grpc.WithTransportCredentials(creds))
	// } else {
		options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// }

	client, _error := grpc.Dial(uri, options...)
	if _error != nil {
		fmt.Printf("Failed to connect to server: %v\n", _error)
		panic(_error)
	}

	return client
}

// 
	// 
	// 

	// 

	// defer client.Close()

	// core.HandleCLI(client, os.Args[1:])