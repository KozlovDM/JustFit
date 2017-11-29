package JSONResponse

import (
	"JustFit/BD"
	"encoding/json"
	"net/http"
	"strconv"
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
	data["status"] = message
	json, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(json)
	//fmt.Fprintf(w, "{%q: %d, %q: %q}", "status", code, "message", message)
}

func ResponseInfo(w http.ResponseWriter, user WorkWithBD.User, subscriptions int, subscribers int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	data := make(map[string]string)
	data["login"] = user.Login
	data["fullname"] = user.Name
	data["publications"] = strconv.Itoa(user.Publication)
	data["info"] = user.Info
	data["subscriptions"] = strconv.Itoa(subscriptions)
	data["subscribers"] = strconv.Itoa(subscribers)
	json, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(json)
	//fmt.Fprintf(w, "{%q: %d, %q: %q}", "status", code, "message", message)
}
