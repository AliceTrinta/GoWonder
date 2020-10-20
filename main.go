package main

import (
	"GoWonder/conf/mongo/dao"
	"GoWonder/controller"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)

func init()  {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization: []string{"Basic Generator"},
		OrganizationalUnit: []string{"Basic generator"},
		CommonName: "Go Web Programming",
	}
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: subject,
		NotBefore: time.Now(),
		NotAfter: time.Now().Add(365 * 24 * time.Hour),
		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
	}
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()

	//Connect to your MongoDatabase
	dao.MongoConnect()
	fmt.Println("Mongo connected")
}

func main()  {
	handler := controller.Begin()
	svr := &http.Server{
		Addr: "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
		Handler: handler,
	}
	log.Fatal(svr.ListenAndServe())
}