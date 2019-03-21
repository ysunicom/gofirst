package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.Get("/viders/:vid-id", streamHandler)
	router.Post("/upload/:vid-id", uploadHandler)
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":9000", r)
}
