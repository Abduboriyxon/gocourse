package main

import (
	"crypto/tls"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	routers "restapi/internal/api/router"

	"github.com/joho/godotenv"
)

//go:embed .env
var envFile embed.FS

func loadEnvFromEmbeddedFile() {
	// Read the embedded .env file
	content, err := envFile.ReadFile(".env")
	if err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	// Create a temp file to load the env vars
	tempFile, err := os.CreateTemp("", ".env")
	if err != nil {
		log.Fatalf("Error creating temp .env file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write content of the embedded .env file to the temp file
	_, err = tempFile.Write(content)
	if err != nil {
		log.Fatalf("Error writing temp .env file: %v", err)
	}

	err = tempFile.Close()
	if err != nil {
		log.Fatalf("Error closing temp file: %v", err)
	}

	// Load env vars from the temp file
	err = godotenv.Load(tempFile.Name())
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
func main() {
	// Only in development, for running secure code
	// err := godotenv.Load()
	// if err != nil {
	// 	return
	// }

	// Load envirnoment variables from the embedded .env file
	loadEnvFromEmbeddedFile()

	// fmt.Println("envirnoment variables CERT_FILE:", os.Getenv("CERT_FILE"))

	port := os.Getenv("API_PORT")

	// cert := "cert.pem"
	// key := "key.pem"

	// cert := os.Getenv("CERT_FILE")
	// key := os.Getenv("KEY_FILE")

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// rl := mw.NewRateLimiter(5, time.Minute)

	// hppOptions := mw.HPPOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// secureMux := mw.Cors(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))
	// secureMux := jwtMiddleware(mw.SecurityHeaders(router))
	// secureMux := (mw.SecurityHeaders(router))
	// secureMux := mw.XSSMiddleware(router)
	router := routers.MainRouter()
	// jwtMiddleware := mw.MiddlewaresExcludePaths(mw.JWTMiddleware, "/execs/login", "/execs/forgotpassword", "/execs/resetpassword/reset")
	// secureMux := utils.ApplyMiddlewares(router, mw.SecurityHeaders, mw.Compression, mw.Hpp(hppOptions), mw.XSSMiddleware, jwtMiddleware, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
	// secureMux := utils.ApplyMiddlewares(router, mw.SecurityHeaders, mw.Compression, mw.Hpp(hppOptions), mw.XSSMiddleware, jwtMiddleware, mw.ResponseTimeMiddleware, mw.Cors)

	// create custom srever
	server := &http.Server{
		Addr: port,
		// Handler: mux,
		Handler: router,
		// Handler: mw.Cors(mux.ServeHTTP),
		TLSConfig: tlsConfig,
	}

	fmt.Println("Starting server on port", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
