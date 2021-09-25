package api

import (
	"encoding/json"
	"fmt"
	"github.com/kayes-shawon/knote/db"
	"github.com/kayes-shawon/knote/pkg/models"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateNote(w http.ResponseWriter, r *http.Request)  {
	note := models.Note{}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &note)
	if err != nil {
		log.Fatalln(err)
	}

	dbCon := db.ConnectDB()

	_, err = dbCon.Model(&note).Insert()
	if err != nil {
		log.Fatalln(err)
	}

	//err = c.JSON(student)



	fmt.Println(note.Tag)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
