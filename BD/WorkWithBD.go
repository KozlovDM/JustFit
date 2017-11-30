package WorkWithBD

import (
	"errors"
	"io"
	"log"
	"mime/multipart"
	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Users
type User struct {
	Name         string
	Login        string
	Phone        string
	Publication  int
	Info         string
	HashPassword []byte
}

type Subscribers struct {
	Login      string
	Subscriber string
}

type Likes struct {
	Image string
	Login string
}

type Comments struct {
	Image   string
	Login   string
	Comment string
}

//SessionMongo
var SessionMongo *mgo.Session

func DeleteData(NameCollection string) error {
	stream := SessionMongo.DB("JustFit").C(NameCollection)
	err := stream.DropCollection()
	return err
}

func DeleteUser(NameDeleted string) error {
	stream := SessionMongo.DB("JustFit").C("Users")
	err := stream.Remove(bson.M{"name": NameDeleted})
	return err
}

func DeleteSub(login string, subscriber string) error {
	stream := SessionMongo.DB("JustFit").C("Subscribers")
	err := stream.Remove(bson.M{"login": login, "subscriber": subscriber})
	return err
}

//AddUser
func AddUser(fullname string, login string, phone string, password []byte) error {
	stream := SessionMongo.DB("JustFit").C("Users")
	err := stream.Insert(&User{Name: fullname, Login: login, Phone: phone, HashPassword: password, Publication: 0})
	return err
}

func UpdateUser(phone string, phonenew string, fullname string, login string, password []byte, info string) error {
	stream := SessionMongo.DB("JustFit").C("Users")
	colQuerier := bson.M{"phone": phone}
	change := bson.M{"$set": bson.M{"info": info, "name": fullname, "login": login, "phone": phonenew, "hashpassword": password}}
	err := stream.Update(colQuerier, change)
	return err
}

func NewSub(login string, subscriber string) error {
	stream := SessionMongo.DB("JustFit").C("Subscribers")
	err := stream.Insert(&Subscribers{Login: login, Subscriber: subscriber})
	return err
}

func NewLike(imageName string, login string) error {
	stream := SessionMongo.DB("JustFit").C("Likes")
	err := stream.Insert(&Likes{Image: imageName, Login: login})
	return err
}

func NewComments(imageName string, login string, comment string) error {
	stream := SessionMongo.DB("JustFit").C("Comments")
	err := stream.Insert(&Comments{Image: imageName, Login: login, Comment: comment})
	return err
}

func FindSub(login string) (result []Subscribers, err error) {
	stream := SessionMongo.DB("JustFit").C("Subscribers")
	err = stream.Find(bson.M{"login": login}).All(&result)
	return result, err
}

func FindSubscriptions(login string) (result []Subscribers, err error) {
	stream := SessionMongo.DB("JustFit").C("Subscribers")
	err = stream.Find(bson.M{"subscriber": login}).All(&result)
	return result, err
}

func FindLikes(imageName string) (result []Likes, err error) {
	stream := SessionMongo.DB("JustFit").C("Likes")
	err = stream.Find(bson.M{"image": imageName}).All(&result)
	return result, err
}

func DeleteLike(imageName string, login string) error {
	stream := SessionMongo.DB("JustFit").C("Likes")
	err := stream.Remove(bson.M{"image": imageName, "login": login})
	return err
}

func FindComments(imageName string) (result []Comments, err error) {
	stream := SessionMongo.DB("JustFit").C("Comments")
	err = stream.Find(bson.M{"image": imageName}).All(&result)
	return result, err
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

func NewPublication(login string) error {
	stream := SessionMongo.DB("JustFit").C("Users")
	user := User{}
	err := stream.Find(bson.M{"login": login}).One(&user)
	if err != nil {
		return err
	}
	count := user.Publication + 1
	colQuerier := bson.M{"login": login}
	change := bson.M{"$set": bson.M{"publication": count}}
	err = stream.Update(colQuerier, change)
	return err
}

func UploadFile(f *multipart.FileHeader, NameCollection string) (interface{}, error) {
	user := User{}
	err := SessionMongo.DB("JustFit").C("Users").Find(bson.M{"login": NameCollection}).One(&user)
	if err != nil {
		return nil, err
	}
	count := strconv.Itoa(user.Publication)
	NameFile := NameCollection + "file" + count
	db := SessionMongo.DB("JustFit")
	file, err := f.Open()
	if err != nil {
		return nil, err
	}
	GridFile, err := db.GridFS(NameCollection).Create(NameFile)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(GridFile, file)
	if err != nil {
		return nil, err
	}
	err = GridFile.Close()
	if err != nil {
		return nil, err
	}
	return GridFile.Id(), nil
}

func DeleteAvatar(login string) error {
	return SessionMongo.DB("JustFit").GridFS("Avatar").Remove(login)
}

func UploadAvatar(f *multipart.FileHeader, login string) (interface{}, error) {
	db := SessionMongo.DB("JustFit")
	file, err := f.Open()
	if err != nil {
		return nil, err
	}
	GridFile, err := db.GridFS("Avatar").Create(login)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(GridFile, file)
	if err != nil {
		return nil, err
	}
	err = GridFile.Close()
	if err != nil {
		return nil, err
	}
	return GridFile.Id(), nil
}

func GetAvatar(login string) ([]byte, error) {
	stream := SessionMongo.DB("JustFit").GridFS("Avatar")
	iter := stream.Find(bson.M{"filename": login}).Iter()

	var image *mgo.GridFile
	if stream.OpenNext(iter, &image) {
		b := make([]byte, image.Size())
		_, err := image.Read(b)
		if err != nil {
			return nil, err
		}
		return b, nil
	}
	err := errors.New("Not Found")
	return nil, err
}

func GetFile(Name string, id interface{}) ([]byte, error) {
	file, err := SessionMongo.DB("JustFit").GridFS(Name).OpenId(id)
	if err != nil {
		return nil, err
	}
	b := make([]byte, file.Size())
	_, err = file.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GetFiles(Name string) (map[string]interface{}, int, error) {
	stream := SessionMongo.DB("JustFit").GridFS(Name)
	iter := stream.Find(nil).Iter()

	result := make(map[string]interface{})
	var image *mgo.GridFile

	i := 0
	var name string
	for stream.OpenNext(iter, &image) {
		i++
		name = "file" + strconv.Itoa(i)
		b := make([]byte, image.Size())
		_, err := image.Read(b)
		if err != nil {
			return nil, i, err
		}
		result[name] = b
	}
	return result, i, iter.Err()
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
