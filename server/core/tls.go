package core

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc/credentials"
)

func GetTLSCredentials() credentials.TransportCredentials {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("cert/server.crt", "cert/server.pem")
	if err != nil {
		fmt.Println("Failed to load server's certificate and private key")
		panic(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config)
}
