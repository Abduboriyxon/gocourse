package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
)

func rootHandler (w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello root route!"))
	fmt.Println("Received request for root route")
}

func teachersHandler (w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on teachers route!"))
		fmt.Println("Received GET Method request for teachers route!")
		return
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on teachers route!"))
		fmt.Println("Received POST Method request for teachers route!")
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on teachers route!"))
		fmt.Println("Received put Method request for teachers route!")
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on teachers route!"))
		fmt.Println("Received PATCH Method request for teachers route!")
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on teachers route!"))
		fmt.Println("Received DELETE Method request for teachers route!")
		return
	}
}

func studentsHandler (w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on students route!"))
		fmt.Println("Received GET Method request for students route!")
		return
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on students route!"))
		fmt.Println("Received POST Method request for students route!")
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on students route!"))
		fmt.Println("Received put Method request for students route!")
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on students route!"))
		fmt.Println("Received PATCH Method request for students route!")
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on students route!"))
		fmt.Println("Received DELETE Method request for students route!")
		return
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on execs route!"))
		fmt.Println("Received GET Method request for execs route!")
		return
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on execs route!"))
		fmt.Println("Received POST Method request for execs route!")
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on execs route!"))
		fmt.Println("Received put Method request for execs route!")
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on execs route!"))
		fmt.Println("Received PATCH Method request for execs route!")
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on execs route!"))
		fmt.Println("Received DELETE Method request for execs route!")
		return
	}
}

func main() {

	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/students/", studentsHandler)

	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr: 	port,
		// Handler: mux,
		Handler: mw.Compression(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Cors(mux)))),
		// Handler: mw.Cors(mux.ServeHTTP),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Starting server on port", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}