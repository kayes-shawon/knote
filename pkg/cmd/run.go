package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/kayes-shawon/knote/pkg/models"
	_ "github.com/kayes-shawon/knote/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request)  {
	Note := models.Note{}
	//err :=
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &Note)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(Note.Tag)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func RunServer() {
	http.HandleFunc("/create/post", createHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
