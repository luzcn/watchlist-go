package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/luzcn/watchlist-go/src/db"
	"net/http"
)

func health(res http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(res, "{}\n")
}

// curl -X POST \
//  http://localhost:5000/notes \
//  -H 'Content-Type: application/json' \
//  -d '{"Note":"abc", "Context":"123"}'
func createNotes(env *db.Env) http.HandlerFunc {

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// decoder := json.Decoder{}
		notes := db.Notes{}
		err := json.NewDecoder(req.Body).Decode(&notes)
		if err != nil {
			panic(err)
		}
		env.AddNote(&notes)
		_, _ = fmt.Fprintf(res, notes.Note)
	})
}

func Load(router *mux.Router, env *db.Env) {
	router.HandleFunc("/health", health).Methods("GET")
	router.HandleFunc("/notes", createNotes(env)).Methods("POST")
}
