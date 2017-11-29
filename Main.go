package main

import (
	"JustFit/Authorization"
	"JustFit/File"
	"JustFit/User"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/SignUp", Authorization.SingUp)
	http.HandleFunc("/SignIn", Authorization.SingIn)
	http.HandleFunc("/Upload", File.Upload)
	http.HandleFunc("/Download", File.Download)
	http.HandleFunc("/UploadAvatar", File.UploadAvatar)
	http.HandleFunc("/GetUserData", User.UserInfo)
	http.HandleFunc("/Sub", User.Subscribe)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
