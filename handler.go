package main

import "net/http"

func users(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"username":"pallat", "name":"pallat","lastname":"anchaleechamaikorn"}`))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}
