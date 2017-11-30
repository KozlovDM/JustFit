package User

import (
	"JustFit/BD"
	"JustFit/File"
	"JustFit/JSONResponse"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func Subscribe(write http.ResponseWriter, request *http.Request) {
	subscriber := request.PostFormValue("subscriber")
	login := request.PostFormValue("login")
	if subscriber == login {
		JSONResponse.ResponseWhithMessage(write, "Неверный данные", http.StatusBadRequest)
		return
	}
	if WorkWithBD.IsLoginExist(login) && WorkWithBD.IsLoginExist(subscriber) {
		sub, err := WorkWithBD.FindSub(login)
		for _, v := range sub {
			if v.Subscriber == subscriber {
				JSONResponse.ResponseWhithMessage(write, "Вы уже подписаны", http.StatusBadRequest)
				return
			}
		}
		err = WorkWithBD.NewSub(login, subscriber)
		if err != nil {
			JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
			return
		}
		JSONResponse.ResponseWhithMessage(write, "Вы подписаны", http.StatusOK)
		return
	}
	JSONResponse.ResponseWhithMessage(write, "Неверный данные", http.StatusBadRequest)
	return
}

func GetUserData(write http.ResponseWriter, request *http.Request, user WorkWithBD.User) {
	subscriptions := 0
	subscribers := 0
	sub, err := WorkWithBD.FindSub(user.Login)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	for _ = range sub {
		subscribers++
	}

	sub, err = WorkWithBD.FindSubscriptions(user.Login)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	for _ = range sub {
		subscriptions++
	}
	result := File.Download(user.Login)
	if result == nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	result["login"] = user.Login
	result["fullname"] = user.Name
	result["publications"] = strconv.Itoa(user.Publication)
	result["info"] = user.Info
	result["subscriptions"] = strconv.Itoa(subscriptions)
	result["subscribers"] = strconv.Itoa(subscribers)
	JSONResponse.ResponseWhithAllData(write, result, http.StatusOK)
}

func UserInfo(write http.ResponseWriter, request *http.Request) {
	phone := request.PostFormValue("phone")
	if !WorkWithBD.IsPhoneExist(phone) {
		JSONResponse.ResponseWhithMessage(write, "Неверный данные", http.StatusBadRequest)
		return
	}
	user := WorkWithBD.FindUserPhone(phone)
	GetUserData(write, request, user)
}

func Like(write http.ResponseWriter, request *http.Request) {
	phone := request.PostFormValue("phone")
	filename := request.PostFormValue("filename")
	if !WorkWithBD.IsPhoneExist(phone) {
		JSONResponse.ResponseWhithMessage(write, "Неверный данные", http.StatusBadRequest)
		return
	}

	user := WorkWithBD.FindUserPhone(phone)
	err := WorkWithBD.DeleteLike(filename, user.Login)
	if err == nil {
		JSONResponse.ResponseWhithMessage(write, "Вам больше не нравится", http.StatusOK)
		return
	}
	err = WorkWithBD.NewLike(filename, user.Login)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	JSONResponse.ResponseWhithMessage(write, "Вам нравится", http.StatusOK)
}

func Сomment(write http.ResponseWriter, request *http.Request) {
	phone := request.PostFormValue("phone")
	filename := request.PostFormValue("filename")
	comment := request.PostFormValue("comment")
	if !WorkWithBD.IsPhoneExist(phone) {
		JSONResponse.ResponseWhithMessage(write, "Неверный данные", http.StatusBadRequest)
		return
	}
	if filename != "" && comment != "" {
		JSONResponse.ResponseWhithMessage(write, "Неверный данные", http.StatusBadRequest)
		return
	}

	user := WorkWithBD.FindUserPhone(phone)
	err := WorkWithBD.NewComments(filename, user.Login, comment)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	JSONResponse.ResponseWhithMessage(write, "Новый комментарий", http.StatusOK)
}

func FindUser(write http.ResponseWriter, request *http.Request) {
	login := request.PostFormValue("login")

	if !WorkWithBD.IsLoginExist(login) {
		JSONResponse.ResponseWhithMessage(write, "Неверный данные", http.StatusBadRequest)
		return
	}
	user := WorkWithBD.FindUserLogin(login)
	GetUserData(write, request, user)
}

func UpdateInfo(write http.ResponseWriter, request *http.Request) {
	phone := request.PostFormValue("phone")
	if !WorkWithBD.IsPhoneExist(phone) {
		JSONResponse.ResponseWhithMessage(write, "Неверный данные", http.StatusBadRequest)
		return
	}

	phonenew := request.FormValue("phonenew")
	fullname := request.FormValue("fullname")
	login := request.FormValue("login")
	password := request.FormValue("password")
	info := request.FormValue("info")
	if phone == "" || fullname == "" || login == "" || password == "" {
		JSONResponse.ResponseWhithMessage(write, "Неккоректные данные", http.StatusBadRequest)
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	err = WorkWithBD.UpdateUser(phone, phonenew, fullname, login, hashPassword, info)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}

	_, avatar, err := request.FormFile("avatar")
	if avatar != nil && err == nil {
		err = WorkWithBD.DeleteAvatar(login)
		if err != nil {
			JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
			return
		}
		_, err = WorkWithBD.UploadAvatar(avatar, login)
		if err != nil {
			JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
			return
		}
	}
	JSONResponse.ResponseWhithMessage(write, "Данные обновлены", http.StatusOK)
}
