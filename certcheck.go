package main

import (
	"crypto/tls"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s <domain name>\n", os.Args[0])
		return
	}

	host := os.Args[1]
	connection, err := tls.Dial("tcp", host+":443", nil)
	if err != nil {
		panic(err)
	}

	certs := connection.ConnectionState().PeerCertificates

	fmt.Printf("Created %s\n", certs[0].NotBefore)
	fmt.Printf("Expires %s\n", certs[0].NotAfter)
	fmt.Printf("Issued by: %s\n", certs[0].Issuer)
}
