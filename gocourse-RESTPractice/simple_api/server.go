package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func loadClientCAs() *x509.CertPool {
	clientCAs := x509.NewCertPool()
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Fatal("Could not read CA certificate:", err)
	}
	clientCAs.AppendCertsFromPEM(caCert)
	return clientCAs
}

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling incoming orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling users")
	})

	port := 3000

	// Load the TLS cert and key
	cert := "cert.pem"
	key := "key.pem"

	// Configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ClientAuth: tls.RequireAndVerifyClientCert, // enforce mTLS
		ClientCAs:  loadClientCAs(),
	}

	// Create a custom server
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		Handler: nil,
		TLSConfig: tlsConfig,
	}

	// Enble http2
	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Starting server on port:",port)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatal("Could not start server:", err)
	}

	// HTTP 1.1 server without TLS
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func logRequestDetails(r *http.Request) {
	httpVersion := r.Proto
	fmt.Println("Recieved request with HTTP Version:", httpVersion)

	if r.TLS != nil {
		tlsVersion := getTLSVersionName(r.TLS.Version)
		fmt.Println("Recieved request with TLS Version:", tlsVersion)
	} else {
		fmt.Println("Recieved request without TLS")
	}
}

func getTLSVersionName(version uint16) string {
	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "Unknown TLS Version"
	}
}