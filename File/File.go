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

	}

	result := make(map[string]interface{})

	phone := request.FormValue("phone")
	user := WorkWithBD.FindUserPhone(phone)
	if !WorkWithBD.IsLoginExist(user.Login) {
		fmt.Println(err)
		JSONResponse.ResponseWhithMessage(write, "Неккоректные данные", http.StatusBadRequest)
		return
	}
	NameCollection += user.Login
	nameFile, err := WorkWithBD.UploadFile(handler, NameCollection)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	result["nameimage"] = nameFile
	result["publications"], err = WorkWithBD.NewPublication(user.Login)
	if err != nil {
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	result["file"], err = WorkWithBD.GetFile(NameCollection, nameFile)
	if err != nil {
		fmt.Println(err)
		JSONResponse.ResponseWhithMessage(write, "Внутренняя ошибка", http.StatusInternalServerError)
		return
	}
	JSONResponse.ResponseWhithAllData(write, result, http.StatusOK)
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

func Download(login string) map[string]interface{} {
	result, count, err := WorkWithBD.GetFiles(login)
	if err != nil {
		return nil
	}
	avatar, err := WorkWithBD.GetAvatar(login)
	if err != nil {
		result["avatar"] = nil
	} else {
		result["avatar"] = avatar
	}
	result["publications"] = count
	return result
}
