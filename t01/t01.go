package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"time"
)

func main() {
	rawPkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	cert := &x509.Certificate{

		Subject: pkix.Name{
			Country:    []string{"KR"},
			CommonName: "Hello",
		},
		SerialNumber: big.NewInt(0),
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(100 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
	}
	rawCert, err := x509.CreateCertificate(rand.Reader, cert, cert, &rawPkey.PublicKey, rawPkey)
	if err != nil {
		panic(err)
	}

	pemPKCS1 := pem.Block{
		Type: "RSA PRIVATE KEY",
		// Headers: map[string]string{},
		Bytes: x509.MarshalPKCS1PrivateKey(rawPkey),
	}
	pemCERT := pem.Block{
		Type: "CERTIFICATE",
		// Headers: map[string]string{},
		Bytes: rawCert,
	}
	fmt.Println("==================")
	fmt.Println(string(pem.EncodeToMemory(&pemCERT)))
	fmt.Println("==================")
	fmt.Println(string(pem.EncodeToMemory(&pemPKCS1)))
	fmt.Println()
	fmt.Println()
	fmt.Println()

	tcpconn, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	tlscfg, err := tls.X509KeyPair(pem.EncodeToMemory(&pemCERT), pem.EncodeToMemory(&pemPKCS1))
	if err != nil {
		panic(err)
	}
	tlsconn := tls.NewListener(tcpconn, &tls.Config{
		Certificates: []tls.Certificate{
			tlscfg,
		},
	})
	http.Serve(tlsconn, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		rw.Write([]byte("Hello, World!"))
		rw.WriteHeader(http.StatusOK)
	}))
}
