package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func rootHandler (w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello root route!")
	w.Write([]byte("Hello root route!"))
	fmt.Println("Received request for root route")
}

func teachersHandler (w http.ResponseWriter, r *http.Request){
	// teachers/{id}
	// teachers/?key=value&query=value2&sortby=email&sortorder=ASC
	switch r.Method {
	case http.MethodGet:
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimSuffix(path, "/")

		fmt.Println("The ID:", userID)

		fmt.Println("Query params",r.URL.Query())
		queryParams := r.URL.Query()
		sortBy := queryParams.Get("sortby")
		key := queryParams.Get("key")
		sortOrder := queryParams.Get("sortorder")

		fmt.Printf("Sort by: %s, Key: %s, Sort order: %s\n", sortBy, key, sortOrder)


		w.Write([]byte("Hello GET Method on teachers route!"))
		// fmt.Println("Received GET Method request for teachers route!")
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

	w.Write([]byte("Hello teachers route!"))
	fmt.Println("Received request for teachers route!")
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
	w.Write([]byte("Hello students route!"))
	fmt.Println("Received request for students route!")
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
	w.Write([]byte("Hello execs route!"))
	fmt.Println("Received requst for execs route!")
}

func main() {

	port := ":3000"

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/teachers/", teachersHandler)

	http.HandleFunc("/students/", studentsHandler)

	http.HandleFunc("/execs/", execsHandler)

	fmt.Println("Starting server on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}