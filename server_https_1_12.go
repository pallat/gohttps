package main

import (
	// "fmt"
	// "io"
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("This is an example server.\n"))
		// fmt.Fprintf(w, "This is an example server.\n")
		// io.WriteString(w, "This is an example server.\n")
	})

	err := s.ListenAndServeTLS("local.crt", "server.key")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
