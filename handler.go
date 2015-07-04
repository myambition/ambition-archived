package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hi there, I loves %s!", r.URL.Path[1:])
}
