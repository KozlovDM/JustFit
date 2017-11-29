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
	_, handler, err := request.FormFile("file")
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
		fmt.Println(err)
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

func UploadAvatar(write http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	_, handler, err := request.FormFile("file")
	if err != nil || handler == nil {
		JSONResponse.ResponseWhithMessage(write, "Неккоректные данные", http.StatusBadRequest)
		return
	}

	phone := request.FormValue("phone")
	user := WorkWithBD.FindUserPhone(phone)
	if !WorkWithBD.IsLoginExist(user.Login) {
		JSONResponse.ResponseWhithMessage(write, "Неккоректные данные", http.StatusBadRequest)
		return
	}
	_, err = WorkWithBD.UploadAvatar(handler, user.Login)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}

	file, err := WorkWithBD.GetAvatar(user.Login)
	if err != nil {
		fmt.Println(err)
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	JSONResponse.ResponseWhithData(write, file, http.StatusOK)
}

func Download(write http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	phone := request.FormValue("phone")
	user := WorkWithBD.FindUserPhone(phone)
	var NameCollection string

	if !WorkWithBD.IsLoginExist(user.Login) {
		JSONResponse.ResponseWhithMessage(write, "Неккоректные данные", http.StatusBadRequest)
		return
	}

	NameCollection += user.Login
	result, count, err := WorkWithBD.GetFiles(NameCollection)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	avatar, err := WorkWithBD.GetAvatar(user.Login)
	if err != nil {
		result["avatar"] = nil
	} else {
		result["avatar"] = avatar
	}
	// publications := make([]byte, 1)
	// publications[0] = byte(count)
	result["publications"] = count
	JSONResponse.ResponseWhithAllData(write, result, http.StatusOK)
}
