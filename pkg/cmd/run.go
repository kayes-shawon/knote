package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/kayes-shawon/knote/pkg/models"
	"log"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request)  {
	note := new(models.Note)

	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "json: %+v", note)
}

func RunServer() {
	http.HandleFunc("/create/post", createHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
