package cmd

import (
	"github.com/kayes-shawon/knote/pkg/api"
	_ "github.com/kayes-shawon/knote/pkg/models"
	"log"
	"net/http"
)



func RunServer() {
	http.HandleFunc("/create/post", api.CreateNote)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
