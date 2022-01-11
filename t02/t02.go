package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

//go:generate openssl genrsa -out priv.pem 2048
//go:generate openssl req -new -x509 -key priv.pem -subj "/C=US/C=KR" -out cert.pem
func main() {
	bts, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		panic(err)
	}
	pem, rest := pem.Decode(bts)
	if len(rest) != 0 {
		panic(rest)
	}

	fmt.Println(pem)
	cert, err := x509.ParseCertificate(pem.Bytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(cert.Subject.Country)
}
