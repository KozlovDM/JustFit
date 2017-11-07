package main

import (
	"JustFit/Authorization"
	"JustFit/File"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/SignUp", Authorization.SingUp)
	http.HandleFunc("/SignIn", Authorization.SingIn)
	http.HandleFunc("/Upload", File.Upload)
	http.HandleFunc("/Download", File.Download)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
