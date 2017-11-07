package File

import (
	"JustFit/BD"
	"JustFit/JSONResponse"
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

	login := request.FormValue("login")
	if !WorkWithBD.IsLoginExist(login) {
		JSONResponse.ResponseWhithMessage(write, "Неккоректные данные", http.StatusBadRequest)
		return
	}
	NameCollection += login
	WorkWithBD.UploadFile(handler, NameCollection)
	JSONResponse.ResponseWhithMessage(write, "Успешная загрузка", http.StatusOK)
}

func Download(write http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	image := request.FormValue("image")
	video := request.FormValue("video")
	login := request.FormValue("login")

	if !WorkWithBD.IsLoginExist(login) {
		JSONResponse.ResponseWhithMessage(write, "Неккоректные данные", http.StatusBadRequest)
		return
	}

	if image == "" || video == "" {
		JSONResponse.ResponseWhithMessage(write, "Неккоректныеданные", http.StatusBadRequest)
		return
	}
	result, err := WorkWithBD.GetFiles(login)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	for _, v := range result {
		JSONResponse.ResponseWhithData(write, v, http.StatusOK)
	}
}
