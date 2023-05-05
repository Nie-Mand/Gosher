package utils

import (
	"fmt"

	"github.com/jcmturner/gokrb5/v8/client"
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/keytab"
	grpc_krb "github.com/jcmturner/grpckrb"
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
	// usesK5 := GetEnv("USES_K5", "1")
	username := GetEnv("WHO", "me")
	keytabFilePath := GetEnv("KEYTAB_PATH", "/etc/krb5.keytab")
	kconfigFilePath := GetEnv("KCONF_PATH", "/etc/krb5.conf")
	realm := GetEnv("KEYTAB_REALM", "GOSHER.COM")
	kcfg, _ := config.Load(kconfigFilePath)
	kt, _ := keytab.Load(keytabFilePath)
	cl := client.NewWithKeytab(username, realm, kt, kcfg)

	interceptor := &grpc_krb.KRBClientInterceptor{
		KRBClient: cl,
	}

	uri := host + ":" + port
	fmt.Println("Connecting to server at: ", uri)
	var options []grpc.DialOption

	// if usesK5 == "1" {
		options = append(options, grpc.WithUnaryInterceptor(interceptor.Unary()))
		options = append(options, grpc.WithStreamInterceptor(interceptor.Stream()))
		options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if true {
	// } else {
	// }

	client, _error := grpc.Dial(uri, options...)
	if _error != nil {
		fmt.Printf("Failed to connect to server: %v\n", _error)
		panic(_error)
	}

	return client
}