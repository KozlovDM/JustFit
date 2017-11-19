package User

import (
	"JustFit/BD"
	"JustFit/JSONResponse"
	"net/http"
)

func Subscribe(write http.ResponseWriter, request *http.Request) {
	subscriber := request.PostFormValue("subscriber")
	login := request.PostFormValue("login")
	if WorkWithBD.IsLoginExist(login) && WorkWithBD.IsLoginExist(subscriber) {
		err := WorkWithBD.NewSub(login, subscriber)
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
