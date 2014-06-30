package shortner

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/pat"
)

type ShortenedPair struct {
	Destination *url.URL
	Hash        string
}

var knownPairs = make(map[string]*ShortenedPair)

func createHandler(w http.ResponseWriter, r *http.Request) {
	dest, err := url.Parse(r.FormValue("url"))
	if err != nil {
		http.Error(w, "invalid url", http.StatusBadRequest)
		return
	}

	h := md5.New()
	io.WriteString(h, dest.String())
	hash := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Printf("%s --> %s\n", hash, dest.String())
	knownPairs[hash] = &ShortenedPair{dest, hash}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(hash))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query().Get(":id")
	match, ok := knownPairs[hash]
	if ok {
		dest := match.Destination
		fmt.Printf("%s --> %s\n", hash, dest.String())
		http.Redirect(w, r, dest.String(), http.StatusFound)
	} else {
		fmt.Printf("%s --> ()\n", hash)
		http.NotFound(w, r)
	}
}

func Main() {
	p := pat.New()
	p.Post("/create", createHandler)
	p.Get("/{id:[a-zA-Z0-9]+}", redirectHandler)
	fmt.Println("Listening on 8080")
	http.Handle("/", p)
	http.ListenAndServe(":8080", nil)
}
