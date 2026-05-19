package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	routers "restapi/internal/api/router"
)

func main() {

	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// rl := mw.NewRateLimiter(5, time.Minute)

	// hppOptions := mw.HPPOptions{
	// 	CheckQuery: true,
	// 	CheckBody: true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist: []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// secureMux := mw.Cors(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))
	// secureMux := utils.ApplyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
	router := routers.Router()
	secureMux := mw.SecurityHeaders(router)

	// create custom srever
	server := &http.Server{
		Addr: 	port,
		// Handler: mux,
		Handler: secureMux,
		// Handler: mw.Cors(mux.ServeHTTP),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Starting server on port", port)
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}