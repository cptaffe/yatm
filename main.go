package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/cptaffe/yatm/matrix"
	"github.com/cptaffe/yatm/matrix/client"
	"github.com/cptaffe/yatm/matrix/client/login"
)

func main() {
	certFile := flag.String("cert", "", "path to TLS cert file")
	keyFile := flag.String("key", "", "path to TLS key file")
	flag.Parse()
	mux := http.NewServeMux()
	mux.Handle("/_matrix/", matrix.New(client.New(login.New())))
	log.Fatal((&http.Server{
		Addr:     ":8080",
		Handler:  mux,
		ErrorLog: log.New(os.Stderr, "logger: ", log.Lshortfile),
	}).ListenAndServeTLS(*certFile, *keyFile))
}
