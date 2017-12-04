package JSONResponse

import (
	"encoding/json"
	"net/http"
)

func ResponseWhithAllData(w http.ResponseWriter, data map[string]interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	js, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(js)
}

func ResponseWhithData(w http.ResponseWriter, jsonData []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	data := make(map[string][]byte)
	data["image"] = jsonData
	js, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(js)
}

func ResponseWhithMessage(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	data := make(map[string]string)
	data["message"] = message
	json, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(json)
}

func ResponseLogin(w http.ResponseWriter, login string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data := make(map[string]string)
	data["login"] = login
	json, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(json)
}
