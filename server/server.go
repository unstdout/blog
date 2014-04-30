package server

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	const INDEX = "/index.html"
	pos := len(r.RequestURI) - len(INDEX)
	if pos > 0 && r.RequestURI[pos:] != INDEX {
		// TODO(jrubin) prevent ...//index.html
		w.Header().Set("Location", r.RequestURI+INDEX)
		w.WriteHeader(http.StatusFound)
		return
	}
	http.NotFound(w, r)
	fmt.Fprint(w, "The opposite of standard.")
}
