package JSONResponse

import (
	"encoding/json"
	"net/http"
)

func ResponseWhithData(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(json)
}

func ResponseWhithMessage(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	data := make(map[string]string)
	data["status"] = message
	json, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(json)
	//fmt.Fprintf(w, "{%q: %d, %q: %q}", "status", code, "message", message)
}
