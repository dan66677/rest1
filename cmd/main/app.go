package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handler(w http.ResponseWriter, r *http.Request, params httprouter.Params)

func main() {
	router := httprouter.New()
	router.GET("/", handler)

	httprouter.
}
