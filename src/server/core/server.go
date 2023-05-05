package core

import (
	"log"
	"net"

	"Nie-Mand/Gosher/server/schemas"
	"Nie-Mand/Gosher/server/services"
	"Nie-Mand/Gosher/server/utils"

	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/jcmturner/gokrb5/v8/service"
	grpc_krb "github.com/jcmturner/grpckrb"
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


	// Kerberos Authentication
	usesK5 := utils.GetEnv("USES_K5", "1")
	keytabFilePath := utils.GetEnv("KEYTAB_PATH", "/etc/krb5.keytab")
	kt, _ := keytab.Load(keytabFilePath)
	interceptor := &grpc_krb.KRBServerInterceptor{
		Settings: service.NewSettings(kt),
	}

	var opts []grpc.ServerOption

	if usesK5 == "1" {
		opts = append(opts, grpc.UnaryInterceptor(interceptor.Unary()))
		opts = append(opts, grpc.StreamInterceptor(interceptor.Stream()))
	}

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
	port := ":" + utils.GetEnv("SERVER_PORT", "50051")

	server, listen := CreateServer(port)

	schemas.RegisterGosherServer(server, &services.ServerStruct{})

	listen()
}

