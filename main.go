package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// HelloWorld for hello world
func HelloWorld() string {
	return "Hello World, golang workshop!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, I love %s!", HelloWorld(), r.URL.Path[1:])
}

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	http.HandleFunc("/", handler)
	log.Println("http server run on " + port + " port")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
