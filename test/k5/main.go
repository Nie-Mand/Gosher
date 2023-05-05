package main

import (
	// "encoding/hex"
	// "log"
	// "os"
	// "time"

	// "github.com/jcmturner/gokrb5/v8/config"
	// "github.com/jcmturner/gokrb5/v8/keytab"
	// "github.com/jcmturner/gokrb5/v8/test/testdata"

	"log"
	"os"
	"time"

	"github.com/jcmturner/gokrb5/v8/client"
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/keytab"
)

const (
	kRB5CONF = `[libdefaults]
	default_realm = GOSHER.COM

[realms]
	GOSHER.COM = {
		kdc_ports = 88,750
		kadmind_port = 749
		kdc = localhost
		admin_server = localhost
	}
 `
)

func main() {
	l := log.New(os.Stderr, "GOKRB5 Client: ", log.LstdFlags)

	// //defer profile.Start(profile.TraceProfile).Stop()
	// // Load the keytab
	keytabFilePath := "consumer.keytab"
	keytabBytes, _ := os.ReadFile(keytabFilePath)
	kt := keytab.New()
	err := kt.Unmarshal(keytabBytes)
	if err != nil {
		l.Fatalf("could not load client keytab: %v", err)
	}

	// Load the client krb5 config
	conf, err := config.NewFromString(kRB5CONF)
	if err != nil {
		l.Fatalf("could not load krb5.conf: %v", err)
	}

	addr := GetEnv("KDC_ADDR", "localhost")
	port := GetEnv("KDC_PORT", ":88")
	
	if addr != "" {
		conf.Realms[0].KDC = []string{addr + port}
	}

	// // Create the client with the keytab
	cl := client.NewWithKeytab("gosher-consumer", "GOSHER.COM", kt, conf, client.Logger(l), client.DisablePAFXFAST(true))

	// // Log in the client
	err = cl.Login()
	if err != nil {
		l.Fatalf("could not login client: %v", err)
	}

	for {
		_, _, err := cl.GetServiceTicket("gosher-serv")
		if err != nil {
			l.Printf("failed to get service ticket: %v\n", err)
		}
		time.Sleep(time.Minute * 5)
	}
}