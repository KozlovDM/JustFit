package WorkWithBD

import (
	"mime/multipart"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestAddUser(t *testing.T) {
	pas, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	err := AddUser("qwerty", "qwerty", "79108922222", pas)
	if err != nil {
		t.Error(err)
	}

	err = DeleteUser("Users", "qwerty")
	if err != nil {
		t.Error(err)
	}
}

func TestFindUserPhone(t *testing.T) {
	pas, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	err := AddUser("qwerty", "qwerty", "79108922222", pas)
	if err != nil {
		t.Error(err)
	}

	res := FindUserPhone("79108922222")
	if res.Login == "" || res.Name == "" || res.Phone == "" || res.HashPassword == nil {
		t.Error("Not found")
	}

	err = DeleteUser("Users", "qwerty")
	if err != nil {
		t.Error(err)
	}
}

func TestFindUserLogin(t *testing.T) {
	pas, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	err := AddUser("qwerty", "qwerty", "79108922222", pas)
	if err != nil {
		t.Error(err)
	}

	res := FindUserLogin("qwerty")
	if res.Login == "" || res.Name == "" || res.Phone == "" || res.HashPassword == nil {
		t.Error("Not found")
	}

	err = DeleteUser("Users", "qwerty")
	if err != nil {
		t.Error(err)
	}
}

func TestIsPhoneExist(t *testing.T) {
	pas, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	err := AddUser("qwerty", "qwerty", "79108922222", pas)
	if err != nil {
		t.Error(err)
	}

	if !IsPhoneExist("79108922222") {
		t.Error("Error find")
	}
	if IsPhoneExist("7910892222") {
		t.Error("Error find")
	}

	err = DeleteUser("Users", "qwerty")
	if err != nil {
		t.Error(err)
	}
}

func TestIsLoginExist(t *testing.T) {
	pas, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	err := AddUser("qwerty", "qwerty", "79108922222", pas)
	if err != nil {
		t.Error(err)
	}

	if !IsLoginExist("qwerty") {
		t.Error("Error find")
	}
	if IsLoginExist("qwert") {
		t.Error("Error find")
	}

	err = DeleteUser("Users", "qwerty")
	if err != nil {
		t.Error(err)
	}
}

func TestUploadEmptyFile(t *testing.T) {
	pas, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	err := AddUser("qwerty", "qwerty", "79108922222", pas)
	if err != nil {
		t.Error(err)
	}

	var d multipart.FileHeader
	d.Filename = "qw.jpg"
	err = UploadFile(&d, "imageqwerty")
	if err == nil {
		t.Error(err)
	}

	err = DeleteUser("Users", "qwerty")
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteDataNotFile(t *testing.T) {
	err := DeleteData("imageUsers")
	if err == nil {
		t.Error("Delete nonexistent file")
	}
}

func TestGetFilesNotFile(t *testing.T) {
	data, err := GetFiles("imageUser")
	if err != nil {
		t.Error(err)
	}
	if data == nil {
		t.Error("Not nil")
	}
}
