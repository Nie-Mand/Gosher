package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/jcmturner/goidentity/v6"
	"github.com/jcmturner/gokrb5/v8/client"
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/keytab"
	"github.com/jcmturner/gokrb5/v8/service"
	"github.com/jcmturner/gokrb5/v8/spnego"
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
	l := log.New(os.Stderr, "GOKRB5 Service Tests: ", log.Ldate|log.Ltime|log.Lshortfile)

	// //defer profile.Start(profile.TraceProfile).Stop()
	// // Load the keytab
	keytabFilePath := "serv.keytab"
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


	th := http.HandlerFunc(testAppHandler)
	sv := httptest.NewServer(spnego.SPNEGOKRB5Authenticate(th, kt, service.Logger(l)))

	// // Create the client with the keytab
	cl := client.NewWithKeytab("gosher-serv", "GOSHER.COM", kt, conf, client.Logger(l), client.DisablePAFXFAST(true))

	// // Log in the client
	err = cl.Login()
	if err != nil {
		l.Fatalf("could not login server: %v", err)
	}

}

func testAppHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Fprint(w, "<html>\n<p><h1>TEST.GOKRB5 Handler</h1></p>\n")
	if validuser, ok := ctx.Value(spnego.CTXKeyAuthenticated).(bool); ok && validuser {
		if creds, ok := ctx.Value(spnego.CTXKeyCredentials).(goidentity.Identity); ok {
			fmt.Fprintf(w, "<ul><li>Authenticed user: %s</li>\n", creds.UserName())
			fmt.Fprintf(w, "<li>User's realm: %s</li></ul>\n", creds.Domain())
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Authentication failed")
	}
	fmt.Fprint(w, "</html>")
	return
}