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

	if phone == "" || fullname == "" || login == "" || password == "" {
		JSONResponse.ResponseWhithMessage(write, "Неккоректные данные", http.StatusBadRequest)
		return
	}

	if WorkWithBD.IsPhoneExist(phone) {
<<<<<<< HEAD
<<<<<<< HEAD
		JSONResponse.ResponseWhithMessage(write, "Такой номер телефона уже заригестрирован", http.StatusConflict)
=======
		JSONResponse.ResponseWhithMessage(write, "Такой номер телефона уже зарегестрирован", http.StatusConflict)
>>>>>>> c039ae87024e99b2581efc1c8627aa042410cded
=======
		JSONResponse.ResponseWhithMessage(write, "Такой номер телефона уже зарегестрирован", http.StatusConflict)
>>>>>>> c039ae87024e99b2581efc1c8627aa042410cded
		return
	}

	if WorkWithBD.IsLoginExist(phone) {
		JSONResponse.ResponseWhithMessage(write, "Такой логин уже существует", http.StatusConflict)
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	//Проверка Работоспособности телефона

	err = WorkWithBD.AddUser(fullname, login, phone, hashPassword)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
<<<<<<< HEAD
<<<<<<< HEAD
	JSONResponse.ResponseWhithMessage(write, "Успешная регистрация", http.StatusOK)
=======
	JSONResponse.ResponseWhithMessage(write, "Успешная регестрация", http.StatusOK)
>>>>>>> c039ae87024e99b2581efc1c8627aa042410cded
=======
	JSONResponse.ResponseWhithMessage(write, "Успешная регестрация", http.StatusOK)
>>>>>>> c039ae87024e99b2581efc1c8627aa042410cded
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
<<<<<<< HEAD
<<<<<<< HEAD
		JSONResponse.ResponseWhithMessage(write, "Неверный номер телефона и/или пароль", http.StatusUnauthorized)
=======
		JSONResponse.ResponseWhithMessage(write, "Неверный номер телефона и/или пароль", http.StatusConflict)
>>>>>>> c039ae87024e99b2581efc1c8627aa042410cded
=======
		JSONResponse.ResponseWhithMessage(write, "Неверный номер телефона и/или пароль", http.StatusConflict)
>>>>>>> c039ae87024e99b2581efc1c8627aa042410cded
		return
	}

	if bcrypt.CompareHashAndPassword(result.HashPassword, []byte(password)) != nil {
<<<<<<< HEAD
<<<<<<< HEAD
		JSONResponse.ResponseWhithMessage(write, "Неверный номер телефона и/или пароль", http.StatusUnauthorized)
		return
=======
		JSONResponse.ResponseWhithMessage(write, "Неверный номер телефона и/или пароль", http.StatusConflict)
>>>>>>> c039ae87024e99b2581efc1c8627aa042410cded
=======
		JSONResponse.ResponseWhithMessage(write, "Неверный номер телефона и/или пароль", http.StatusConflict)
>>>>>>> c039ae87024e99b2581efc1c8627aa042410cded
	}
	JSONResponse.ResponseWhithMessage(write, "Успешная авторизация", http.StatusOK)
}
