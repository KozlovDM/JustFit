package Upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request) string {
	r.ParseForm()
	file, handler, err := r.FormFile("image")
	if err != nil {
		//ошибка
	}
	defer file.Close()
	_, err = fmt.Fprintf(w, "%v", handler.Header)
	if err != nil {
		//ошибка
	}
	path := "C:/work/src/" + handler.Filename
	newFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(newFile, file)
	return path
}
