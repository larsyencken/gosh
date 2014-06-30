package shortner

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ShortenedPair struct {
	Destination string
	Hash        string
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func Main() {
	r := mux.NewRouter()
	r.HandleFunc("/create", createHandler)
	r.HandleFunc("/{id:[a-zA-Z0-9]+}", redirectHandler)
	fmt.Println("Listening on 8080")
	http.ListenAndServe(":8080", nil)
}
