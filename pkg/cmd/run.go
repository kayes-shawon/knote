package cmd

import (
	_ "github.com/kayes-shawon/knote/pkg/models"
	"github.com/kayes-shawon/knote/pkg/urls"
	"log"
	"net/http"
)



func RunServer() {
	urls.UrlPatterns()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
