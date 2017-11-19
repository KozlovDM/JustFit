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
	err := stream.Insert(&User{Name: fullname, Login: login, Phone: phone, HashPassword: password})
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

func FindLikes(imageName string) (result []Likes, err error) {
	stream := SessionMongo.DB("JustFit").C("Likes")
	err = stream.Find(bson.M{"image": imageName}).All(&result)
	return result, err
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
	//fmt.Println(result.ID)
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
