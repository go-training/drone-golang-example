package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I love %s!", r.URL.Path[1:])
}

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "server port")
	flag.StringVar(&port, "p", "8080", "server port")
	flag.Parse()

	if p, ok := os.LookupEnv("PORT"); ok {
		port = p
	}

	http.HandleFunc("/", handler)
	log.Println("http server run on " + port + " port")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
