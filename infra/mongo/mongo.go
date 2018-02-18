package mongo

import (
	"fmt"
	"log"

	"github.com/RobinBaeckman/hermod-api/domain"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
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

type ProductDB struct {
	Conn *mgo.Database
}

type AuthDB struct {
	Conn *mgo.Database
}

type HomeDB struct {
	Conn *mgo.Database
}

type AdminDB struct {
	Conn *mgo.Database
}

type LoginDB struct {
	Conn *mgo.Database
}

type IndexDB struct {
	Conn *mgo.Database
}

func NewProductDB(db *mgo.Database) *ProductDB {
	return &ProductDB{
		Conn: db,
	}
}

func NewAuthDB(db *mgo.Database) *AuthDB {
	return &AuthDB{
		Conn: db,
	}
}
func NewHomeDB(db *mgo.Database) *HomeDB {
	return &HomeDB{
		Conn: db,
	}
}
func NewAdminDB(db *mgo.Database) *AdminDB {
	return &AdminDB{
		Conn: db,
	}
}
func NewLoginDB(db *mgo.Database) *LoginDB {
	return &LoginDB{
		Conn: db,
	}
}
func NewIndexDB(db *mgo.Database) *IndexDB {
	return &IndexDB{
		Conn: db,
	}
}

func (db *ProductDB) Store(p domain.Product) (domain.Product, error) {
	fmt.Println("######[DB]########")
	fmt.Println(p)
	fmt.Println("#########################")

	p.ID = uuid.NewV4().String()
	err := db.Conn.C("products").Find(bson.M{"id": p.ID}).One(&p)
	if err != nil {
		log.Fatal(err)
	}

	return p, nil
}

func (db *ProductDB) Get(id string) (domain.Product, error) {
	product := domain.Product{}

	err := db.Conn.C("products").Find(bson.M{"id": id}).One(&product)
	if err != nil {
		log.Fatal(err)
	}

	return product, nil
}

func (db *ProductDB) GetAll() ([]*domain.Product, error) {
	products := []*domain.Product{}
	err := db.Conn.C("products").Find(nil).All(&products)
	if err != nil {
		log.Fatal(err)
	}

	return products, nil
}

func (db *AuthDB) Get(email string) (domain.Admin, error) {
	a := domain.Admin{}
	err := db.Conn.C("admins").Find(bson.M{"email": email}).One(&a)
	if err != nil {
		errors.Wrap(err, "Wrong password or username")
	}

	return a, nil
}

func (db *AuthDB) Store(a domain.Admin) error {
	a.ID = uuid.NewV4().String()
	err := db.Conn.C("admins").Find(bson.M{"id": a.ID}).One(&a)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (db *AdminDB) GetSingle(a *domain.Admin) error {

	err := db.Conn.C("admins").Find(bson.M{"email": a.Email}).One(a)
	if err != nil {
		errors.Wrap(err, "Wrong password or username")
	}

	return nil
}

func (db *AdminDB) Store(a *domain.Admin) error {
	a.ID = uuid.NewV4().String()
	err := db.Conn.C("admins").Insert(a)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (db *AdminDB) GetMany(as *[]domain.Admin) error {
	err := db.Conn.C("admins").Find(nil).All(as)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
