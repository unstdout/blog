package server

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/go", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "The opposite of standard.")
}
