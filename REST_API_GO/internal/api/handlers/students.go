package handlers

import "net/http"

func StudentsHandler (w http.ResponseWriter, r *http.Request){
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