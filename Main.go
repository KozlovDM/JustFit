package Main

import (
	"JustFit/Authorization"
	"JustFit/Upload"
	"net/http"
)

func main() {
	http.HandleFunc("/SingUp", Authorization.SingUp)
	http.HandleFunc("/SingIn", Authorization.SingIn)
	http.HandleFunc("/Upload", Upload.Upload)
	err := http.ListenAndServe(":27017", nil)
	if err != nil {
		//ошибка
	}
}
