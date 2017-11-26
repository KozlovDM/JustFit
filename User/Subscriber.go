package User

import (
	"JustFit/BD"
	"JustFit/JSONResponse"
	"net/http"
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

func UserInfo(write http.ResponseWriter, request *http.Request) {
	phone := request.PostFormValue("phone")
	if !WorkWithBD.IsPhoneExist(phone) {
		JSONResponse.ResponseWhithMessage(write, "Неверный данные", http.StatusBadRequest)
		return
	}
	user := WorkWithBD.FindUserPhone(phone)
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
	JSONResponse.ResponseInfo(write, user, subscriptions, subscribers)
}
