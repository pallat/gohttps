package main

import (
	// "fmt"
	// "io"
	"crypto/tls"
	"log"
	"net/http"
)

// this server support only tls version 1.0
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("This is an example server.\n"))
		// fmt.Fprintf(w, "This is an example server.\n")
		// io.WriteString(w, "This is an example server.\n")
	})

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	s := &http.Server{
		Addr:      ":443",
		TLSConfig: cfg,
	}

	err := s.ListenAndServeTLS("local.crt", "server.key")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
