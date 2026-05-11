package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"strconv"
	"strings"
)

type Teacher struct{
	ID int
	FirstName string
	LastName string
	Class string
	Subject string
}

var (
	teachers = make(map[int]Teacher)
	// mutex = &sync.Mutex{}
	nextID = 1
)

// initialize some dummy data
func init() {
	teachers[nextID] = Teacher{
		ID: nextID,
		FirstName: "John",
		LastName: "Doe",
		Class: "10A",
		Subject: "Math",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID: nextID,
		FirstName: "Jane",
		LastName: "Smith",
		Class: "10B",
		Subject: "Algebra",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID: nextID,
		FirstName: "Jane",
		LastName: "Doe",
		Class: "11A",
		Subject: "Biology",
	}
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	fmt.Println(idStr)

	if idStr == "" { 
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")

		teacherList := make([]Teacher, 0, len(teachers))
		for _, teacher := range teachers{
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName){
			teacherList = append(teacherList, teacher)
			}
		}

		response := struct {
			Status string `json:"status"`
			Count int `json:"count"`
			Data []Teacher `json:"data"`
		}{
			Status: "success",
			Count: len(teacherList),
			Data: teacherList,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

	// Handle Path parameter
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	teacher, exists := teachers[id]
	if !exists {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(teacher)
}

func rootHandler (w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello root route!"))
}

func teachersHandler (w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		getTeachersHandler(w, r)
		return
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on teachers route!"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on teachers route!"))
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on teachers route!"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on teachers route!"))
		return
	}
}

func studentsHandler (w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on students route!"))
		return
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on students route!"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on students route!"))
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on students route!"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on students route!"))
		return
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET Method on execs route!"))
		return
	case http.MethodPost:
		w.Write([]byte("Hello POST Method on execs route!"))
		return
	case http.MethodPut:
		w.Write([]byte("Hello PUT Method on execs route!"))
		return
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH Method on execs route!"))
		return
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE Method on execs route!"))
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

	// rl := mw.NewRateLimiter(5, time.Minute)

	// hppOptions := mw.HPPOptions{
	// 	CheckQuery: true,
	// 	CheckBody: true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist: []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// secureMux := mw.Cors(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)(mux))))))
	// secureMux := applyMiddlewares(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
	secureMux := mw.SecurityHeaders(mux)
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

// Middleware is a function that wraps an http.Handler with additional functionality
type Middleware func(http.Handler) http.Handler

func ApplyMiddlewares(handler http.Handler, middlewares ...Middleware) http.Handler{
	for _, middleware := range middlewares{
		handler = middleware(handler)
	}
	return handler
}