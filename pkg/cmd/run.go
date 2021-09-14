package cmd

import (
	"fmt"
	"log"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func RunServer() {
	http.HandleFunc("/create/post", createHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
