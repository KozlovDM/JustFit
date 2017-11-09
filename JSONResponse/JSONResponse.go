package JSONResponse

import (
	"fmt"
	"net/http"
)

func ResponseWhithData(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	//fmt.Fprintf(w, "{ %q: \"%v\"}", "data", json)
	w.Write(json)
}

func ResponseWhithMessage(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{ %q: %q}", "message", message)
}
