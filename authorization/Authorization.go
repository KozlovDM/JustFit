package Authorization

import (
	"JustFit/BD"
	"io"
	"log"
	"net/http"
)

//SingUp
func SingUp(write http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	phone := request.FormValue("phone")
	fullname := request.FormValue("fullname")
	login := request.FormValue("login")
	password := request.FormValue("password")

	if WorkWithBD.IsPhoneExist(phone) {
		io.WriteString(write, "Такой номер телефона уже зарегестрирован")
		return
	}

	if WorkWithBD.IsLoginExist(phone) {
		io.WriteString(write, "Такой логин уже существует")
		return
	}

	var hashPassword string
	//Hash = Функция получения Хэша ключа
	//Проверка Работоспособности телефона

	err := WorkWithBD.AddUser(fullname, login, phone, hashPassword)
	if err != nil {
		return
	}
}

//SingIn
func SingIn(write http.ResponseWriter, request *http.Request) {

	phone := request.PostFormValue("phone")
	login := request.PostFormValue("login")
	password := request.FormValue("password")

	var result WorkWithBD.Users

	if WorkWithBD.IsPhoneExist(phone) {
		result = WorkWithBD.FindUserPhone(phone)
	} else if WorkWithBD.IsLoginExist(login) {
		result = WorkWithBD.FindUserLogin(login)
	} else {
		io.WriteString(write, "Неверный номер телефона и/или пароль")
		return
	}

	var hashPassword string
	//Hash = Функция получения Хэша ключа

	if result.HashPassword != hashPassword {
		io.WriteString(write, "Неверный номер телефона и/или пароль")
	}
}

func init() {
	http.HandleFunc("/SingUp", SingUp)
	http.HandleFunc("/SingIn", SingIn)
	err := http.ListenAndServe(":27017", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
