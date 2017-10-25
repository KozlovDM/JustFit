package WorkWithBD

import (
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

func init() {
	var err error
	SessionMongo, err = mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	//defer SessionMongo.Close()
	SessionMongo.SetMode(mgo.Monotonic, true)
}
