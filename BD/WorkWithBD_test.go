package WorkWithBD

import (
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestAddUser(t *testing.T) {
	pas, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	err := AddUser("qwerty", "qwerty", "79108922222", pas)
	assert.NoError(t, err)

	err = DeleteUser("qwerty")
	assert.NoError(t, err)
}

func TestFindUserPhone(t *testing.T) {
	pas, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	err := AddUser("qwerty", "qwerty", "79108922222", pas)
	assert.NoError(t, err)

	res := FindUserPhone("79108922222")
	if res.Login == "" || res.Name == "" || res.Phone == "" || res.HashPassword == nil {
		t.Error("Not found")
	}

	err = DeleteUser("qwerty")
	assert.NoError(t, err)
}

func TestFindUserLogin(t *testing.T) {
	pas, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	err := AddUser("qwerty", "qwerty", "79108922222", pas)
	assert.NoError(t, err)

	res := FindUserLogin("qwerty")
	if res.Login == "" || res.Name == "" || res.Phone == "" || res.HashPassword == nil {
		t.Error("Not found")
	}

	err = DeleteUser("qwerty")
	assert.NoError(t, err)
}

func TestIsPhoneExist(t *testing.T) {
	pas, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	err := AddUser("qwerty", "qwerty", "79108922222", pas)
	assert.NoError(t, err)

	assert.EqualValues(t, true, IsPhoneExist("79108922222"))
	assert.EqualValues(t, false, IsPhoneExist("7910892222"))

	err = DeleteUser("qwerty")
	assert.NoError(t, err)
}

func TestIsLoginExist(t *testing.T) {
	pas, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	err := AddUser("qwerty", "qwerty", "79108922222", pas)
	assert.NoError(t, err)

	assert.EqualValues(t, true, IsLoginExist("qwerty"))
	assert.EqualValues(t, false, IsLoginExist("qwert"))

	err = DeleteUser("qwerty")
	assert.NoError(t, err)
}

func TestUploadEmptyFile(t *testing.T) {
	pas, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	err := AddUser("qwerty", "qwerty", "79108922222", pas)
	assert.NoError(t, err)

	var d multipart.FileHeader
	d.Filename = "qw.jpg"
	err = UploadFile(&d, "imageqwerty")
	if err == nil {
		t.Error(err)
	}

	err = DeleteUser("qwerty")
	assert.NoError(t, err)
}

func TestDeleteDataNotFile(t *testing.T) {
	err := DeleteData("imageUsers")
	if err == nil {
		t.Error("Delete nonexistent file")
	}
}

func TestGetFilesNotFile(t *testing.T) {
	data, err := GetFiles("imageUser")
	assert.NoError(t, err)
	if data == nil {
		t.Error("Not nil")
	}
}

func TestLike(t *testing.T) {
	err := NewLike("qwerty", "qwerty")
	assert.NoError(t, err)
	err = NewLike("qwerty", "qwerty1")
	assert.NoError(t, err)

	res, err := FindLikes("qwerty")
	var count int
	for _ = range res {
		count++
	}
	assert.Equal(t, 2, count)

	err = SessionMongo.DB("JustFit").C("Likes").DropCollection()
	assert.NoError(t, err)
}

func TestComment(t *testing.T) {
	err := NewComments("qwerty", "qwerty", "Hello")
	assert.NoError(t, err)
	err = NewComments("qwerty", "qwerty1", "Hello1")
	assert.NoError(t, err)

	res, err := FindComments("qwerty")
	var count int
	for _ = range res {
		count++
	}
	assert.Equal(t, 2, count)

	err = SessionMongo.DB("JustFit").C("Comments").DropCollection()
	assert.NoError(t, err)
}

func TestSub(t *testing.T) {
	err := NewSub("qwerty", "Sub")
	assert.NoError(t, err)
	err = NewSub("qwerty", "Sub1")
	assert.NoError(t, err)
	res, err := FindSub("qwerty")
	assert.NoError(t, err)
	var count int
	for _ = range res {
		count++
	}
	assert.Equal(t, 2, count)

	err = DeleteSub("qwerty", "Sub")
	assert.NoError(t, err)
	count = 0
	res, err = FindSub("qwerty")
	assert.NoError(t, err)
	for _ = range res {
		count++
	}
	assert.Equal(t, 1, count)

	err = SessionMongo.DB("JustFit").C("Subscribers").DropCollection()
	assert.NoError(t, err)
}
