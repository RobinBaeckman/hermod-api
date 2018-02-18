package repo

import (
	"fmt"
	"log"

	"github.com/RobinBaeckman/hermod-api/domain"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("hermod")

	DB = c
}

var DB *mgo.Database

type AuthStore struct {
	session *mgo.Session
}

func NewMgoDB() *MgoDB {
	authStore := AuthStore{
		session: c,
	}

	return &mgoDB
}

func (db *MgoDB) Get(email string) (domain.Admin, error) {
	a := domain.Admin{}
	err := db.Conn.C("admins").Find(bson.M{"email": email}).One(&a)
	if err != nil {
		log.Fatal(err)
	}

	return a, nil
}

func (db *MgoDB) Store(a domain.Admin) error {
	fmt.Println("######[DB]########")
	fmt.Println(a)
	fmt.Println("#########################")

	a.ID = uuid.NewV4().String()
	err := db.Conn.C("products").Insert(&a)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
