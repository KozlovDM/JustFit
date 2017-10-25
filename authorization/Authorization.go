package Authorization

import (
	"JustFit/BD"
	"JustFit/JSONResponse"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

//SingUp
func SingUp(write http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	phone := request.FormValue("phone")
	fullname := request.FormValue("fullname")
	login := request.FormValue("login")
	password := request.FormValue("password")

	if WorkWithBD.IsPhoneExist(phone) {
		JSONResponse.ResponseWhithMessage(write, "Такой номер телефона уже зарегестрирован", http.StatusConflict)
		return
	}

	if WorkWithBD.IsLoginExist(phone) {
		JSONResponse.ResponseWhithMessage(write, "Такой логин уже существует", http.StatusConflict)
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		//ошибка
	}
	//Проверка Работоспособности телефона

	err = WorkWithBD.AddUser(fullname, login, phone, hashPassword)
	if err != nil {
		return
	}
	JSONResponse.ResponseWhithMessage(write, "Успешная регестрация", http.StatusOK)
}

//SingIn
func SingIn(write http.ResponseWriter, request *http.Request) {

	phone := request.PostFormValue("phone")
	login := request.PostFormValue("login")
	password := request.FormValue("password")

	var result WorkWithBD.User

	if WorkWithBD.IsPhoneExist(phone) {
		result = WorkWithBD.FindUserPhone(phone)
	} else if WorkWithBD.IsLoginExist(login) {
		result = WorkWithBD.FindUserLogin(login)
	} else {
		JSONResponse.ResponseWhithMessage(write, "Неверный номер телефона и/или пароль", http.StatusConflict)
		return
	}

	if bcrypt.CompareHashAndPassword(result.HashPassword, []byte(password)) != nil {
		JSONResponse.ResponseWhithMessage(write, "Неверный номер телефона и/или пароль", http.StatusConflict)
	}
	JSONResponse.ResponseWhithMessage(write, "Успешная авторизация", http.StatusOK)
}
