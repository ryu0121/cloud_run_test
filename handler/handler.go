package handler

import (
	"encoding/json"
	"net/http"
)

type Hello struct {
	Message string `json:"message"`
}

func Handle(w http.ResponseWriter, _ *http.Request) {
	h := Hello{
		Message: "Hello World",
	}
	enc := json.NewEncoder(w)
	enc.Encode(&h)
}

var Handler = http.HandlerFunc(Handle)
