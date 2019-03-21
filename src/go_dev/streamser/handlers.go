package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println()
}
