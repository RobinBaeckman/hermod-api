package mongo

import (
	"github.com/RobinBaeckman/hermod-api/customerr"
	"github.com/RobinBaeckman/hermod-api/domain"
	"github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductDB struct {
	Conn *mgo.Database
}

type AuthDB struct {
	Conn *mgo.Database
}

type AdminDB struct {
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
func NewAdminDB(db *mgo.Database) *AdminDB {
	return &AdminDB{
		Conn: db,
	}
}

func (db *ProductDB) Persist(p domain.Product) (domain.Product, error) {
	p.ID = uuid.NewV4().String()
	err := db.Conn.C("products").Find(bson.M{"id": p.ID}).One(&p)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (db *ProductDB) Get(id string) (domain.Product, error) {
	p := domain.Product{}

	err := db.Conn.C("products").Find(bson.M{"id": id}).One(&p)
	if err != nil {
		return p, &customerr.App{err, "There is no product with that id.", 404}
	}

	return p, nil
}

func (db *ProductDB) GetAll() ([]*domain.Product, error) {
	p := []*domain.Product{}
	err := db.Conn.C("products").Find(nil).All(&p)
	if err != nil {
		return p, &customerr.App{err, "There is no product with that id.", 404}
	}

	return p, nil
}

func (db *AuthDB) Get(email string) (domain.Admin, error) {
	a := domain.Admin{}
	err := db.Conn.C("admins").Find(bson.M{"email": email}).One(&a)
	if err != nil {
		return a, &customerr.App{err, "Wrong password or username", 404}
	}

	return a, nil
}

func (db *AuthDB) Persist(a domain.Admin) error {
	a.ID = uuid.NewV4().String()
	err := db.Conn.C("admins").Find(bson.M{"id": a.ID}).One(&a)
	if err != nil {
		return err
	}

	return nil
}

func (db *AdminDB) Get(a *domain.Admin) error {

	err := db.Conn.C("admins").Find(bson.M{"id": a.ID}).One(a)
	if err != nil {
		return &customerr.App{err, "There is no admin user with that id.", 404}
	}

	return nil
}

func (db *AdminDB) Persist(a *domain.Admin) error {
	a.ID = uuid.NewV4().String()
	err := db.Conn.C("admins").Insert(a)
	if err != nil {
		return err
	}

	return nil
}

func (db *AdminDB) GetAll(as *[]domain.Admin) error {
	err := db.Conn.C("admins").Find(nil).All(as)
	if err != nil {
		return err
	}

	return nil
}
