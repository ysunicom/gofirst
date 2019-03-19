package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

//respose
type SigneUp struct {
	Sussess   bool   `json:"sussess"`
	SessionId string `json:"session_id"`
}

//data modle
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

type SimpleSession struct {
	Username string //login name
	TTL      int64
}
