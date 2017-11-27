package File

import (
	"JustFit/BD"
	"JustFit/JSONResponse"
	"fmt"
	"net/http"
)

func Upload(write http.ResponseWriter, request *http.Request) {
	var NameCollection string
	request.ParseForm()
	_, handler, err := request.FormFile("image")
	if err != nil || handler == nil {
		_, handler, err = request.FormFile("video")
		if err != nil || handler == nil {
			JSONResponse.ResponseWhithMessage(write, "Неккоректные данные", http.StatusBadRequest)
			return
		}
		NameCollection = "video"

	} else {
		NameCollection = "image"
	}

	phone := request.FormValue("phone")
	user := WorkWithBD.FindUserPhone(phone)
	if !WorkWithBD.IsLoginExist(user.Login) {
		JSONResponse.ResponseWhithMessage(write, "Неккоректные данные", http.StatusBadRequest)
		return
	}
	NameCollection += user.Login
	id, err := WorkWithBD.UploadFile(handler, NameCollection)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	err = WorkWithBD.NewPublication(user.Login)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	file, err := WorkWithBD.GetFile(NameCollection, id)
	if err != nil {
		fmt.Println(err)
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	JSONResponse.ResponseWhithData(write, file, http.StatusOK)
}

func Download(write http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	image := request.FormValue("image")
	video := request.FormValue("video")
	login := request.FormValue("login")
	var NameCollection string

	if !WorkWithBD.IsLoginExist(login) {
		JSONResponse.ResponseWhithMessage(write, "Неккоректные данные", http.StatusBadRequest)
		return
	}

	if image == "" {
		if video == "" {
			JSONResponse.ResponseWhithMessage(write, "Неккоректные данные", http.StatusBadRequest)
			return
		}
		NameCollection = "video"
	} else {
		NameCollection = "image"
	}
	NameCollection += login
	result, err := WorkWithBD.GetFiles(NameCollection)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	for _, v := range result {
		JSONResponse.ResponseWhithData(write, v, http.StatusOK)
	}
}
