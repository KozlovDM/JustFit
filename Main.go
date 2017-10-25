package main

import (
	"JustFit/Authorization"
	"JustFit/Upload"
	"net/http"
)

func main() {
	http.HandleFunc("/SingUp", Authorization.SingUp)
	http.HandleFunc("/SingIn", Authorization.SingIn)
	http.HandleFunc("/Upload", Upload.Upload)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		//ошибка
	}
}
