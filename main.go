package main

import (
	"net/http"
	"flag"
	"log"
	"mime"
)

var (
	port = flag.String("port", "8000", "Listen at port specified.")
	root = flag.String("root", ".", "Location to host.")
)

func main () {
	mime.AddExtensionType(".svg", "image/svg+xml")

	flag.Parse()
	log.Printf("Listening at %s, hosting %s", *port, *root)
	log.Fatal(http.ListenAndServe(":" + *port, http.FileServer(http.Dir(*root))))
}
