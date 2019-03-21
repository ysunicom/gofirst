package main
{
	"io"
	"encoding/json"
	"net/http"
}
func sendErroResponse(w http.ResponseWriter,sc int,errMsg string){
	w.WriteHeader(sc)
	io.WriteString(w,errMsg)
}