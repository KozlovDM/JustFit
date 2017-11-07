package WorkWithBD

import (
	"io"
	"log"
	"mime/multipart"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Users
type User struct {
	Name         string
	Login        string
	Phone        string
	HashPassword []byte
	Store        mgo.GridFS
}

//SessionMongo
var SessionMongo *mgo.Session

//AddUser
func AddUser(fullname string, login string, phone string, password []byte) error {
	stream := SessionMongo.DB("JustFit").C("Users")
	err := stream.Insert(&User{Name: fullname, Login: login, Phone: phone, HashPassword: password})
	return err
}

//FindUserPhone
func FindUserPhone(phone string) User {
	stream := SessionMongo.DB("JustFit").C("Users")
	result := User{}
	_ = stream.Find(bson.M{"phone": phone}).One(&result)
	return result
}

//FindUserLogin
func FindUserLogin(login string) User {
	stream := SessionMongo.DB("JustFit").C("Users")
	result := User{}
	_ = stream.Find(bson.M{"login": login}).One(&result)
	return result
}

//IsPhoneExist
func IsPhoneExist(phone string) bool {
	stream := SessionMongo.DB("JustFit").C("Users")
	user := User{}
	err := stream.Find(bson.M{"phone": phone}).One(&user)
	if err != nil {
		return false
	}
	return true
}

//IsLoginExist
func IsLoginExist(login string) bool {
	stream := SessionMongo.DB("JustFit").C("Users")
	user := User{}
	err := stream.Find(bson.M{"login": login}).One(&user)
	if err != nil {
		return false
	}
	return true
}

func UploadFile(f *multipart.FileHeader, NameCollection string) error {
	NameFile := NameCollection + f.Filename
	db := SessionMongo.DB("JustFit")
	file, err := f.Open()
	if err != nil {
		return err
	}
	GridFile, err := db.GridFS(NameCollection).Create(NameFile)
	if err != nil {
		return err
	}
	_, err = io.Copy(GridFile, file)
	if err != nil {
		return err
	}
	err = GridFile.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetFiles(Name string) (map[string][]byte, error) {
	stream := SessionMongo.DB("JustFit").GridFS(Name)
	iter := stream.Find(nil).Iter()

	result := make(map[string][]byte)
	var image *mgo.GridFile
	for stream.OpenNext(iter, &image) {

		b := make([]byte, image.Size())
		_, err := image.Read(b)
		if err != nil {
			return nil, err
		}
		result[image.Name()] = b
	}
	return result, iter.Err()
}

func init() {
	var err error
	SessionMongo, err = mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatalln(err)
	}
	//defer SessionMongo.Close()
	SessionMongo.SetMode(mgo.Monotonic, true)
}
