package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "github.com/kayes-shawon/knote/pkg/models"
)

func createHandler(w http.ResponseWriter, r *http.Request)  {
	//Note := models.Note{}
	//err :=
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func RunServer() {
	http.HandleFunc("/create/post", createHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
