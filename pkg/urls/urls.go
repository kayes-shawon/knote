package urls

import (
	"github.com/kayes-shawon/knote/pkg/api"
	"net/http"
)

func UrlPatterns() {
	http.HandleFunc("/create/post", api.CreateNote)
}
