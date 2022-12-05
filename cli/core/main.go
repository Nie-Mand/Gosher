package core

import (
	"context"
	"fmt"
	"os"
	"strings"

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

/*
* @function: GetFile
* @description: Open a file and return the file
* @params: filePath: string
* @return: []byte
 */
func GetFile(filePath string) []byte {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error opening file")
		panic(err)
	}
	return file
}

/*
* @function: SaveFile
* @description: Save a file into the downloads directory
* @params: filePath: string, file: []byte
* @return: void
 */
func SaveFile(filePath string, file []byte) {
	err := os.WriteFile(filePath,
		file,
		0644,
	)

	if err != nil {
		fmt.Println("Error saving file")
		panic(err)
	} else {
		fmt.Println("Saved file")
	}
}

/*
* @function: CheckFilesForMatch
* @description: Check folder for files
* @params: description: string
* @return: string[]
 */
func CheckFilesForMatch(description string) []string {
	files, err := os.ReadDir("downloads")
	if err != nil {
		fmt.Println("Error reading directory")
		panic(err)
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() && strings.Contains(file.Name(), description) {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames
}

/*
* @function: OpenFile
* @description: Open a file and return the file as a byte array
* @params: filename: string
* @return: []byte
 */
func OpenFile(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening file")
		fmt.Println(err)
		return nil, err
	}

	return file, nil
}
