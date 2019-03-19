package main

import (
	"encoding/json"
	"go_dev/videoser/api/defs"
	"io"
	"net/http"
)

func sendErrorRespose(w http.ResponseWriter, errResp defs.ErrResponse) {
	w.WriteHeader(errResp.HttpSC)
	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}
func setNormalRespose(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
