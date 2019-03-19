package main

import (
	"encoding/json"
	"go_dev/videoser/api/dbops"
	"go_dev/videoser/api/defs"
	"go_dev/videoser/api/session"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorRespose(w, defs.ErrorRequestBodyParaseFailed)
		return
	}
	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorRespose(w, defs.ErrorDBError)
		return
	}
	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SigneUp{Sussess: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorRespose(w, defs.ErrorInternalFaults)
		return
	} else {
		setNormalRespose(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
