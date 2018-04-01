package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("a", ":10000", "listen address")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir("./src")))

	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
