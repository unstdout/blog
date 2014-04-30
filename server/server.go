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
	l := len(r.RequestURI)
	pos := l - len(INDEX)
	if pos > 0 && r.RequestURI[pos:] != INDEX {
		url := r.RequestURI
		if url[l-1] != '/' {
			url += "/"
		}
		w.Header().Set("Location", url+INDEX[1:])
		w.WriteHeader(http.StatusFound)
		return
	}
	http.NotFound(w, r)
	fmt.Fprint(w, "The opposite of standard.")
}
