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

	s := &http.Server{
		Addr: ":443",
		TLSConfig: &tls.Config{
			// PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_RSA_WITH_RC4_128_SHA,
				tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
				tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			},
			MinVersion: tls.VersionTLS10,
			MaxVersion: tls.VersionTLS10,
		},
	}

	err := s.ListenAndServeTLS("local.crt", "server.key")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
