package product

import (
	"fmt"

	"github.com/RobinBaeckman/hermod-api/domain"
	"github.com/RobinBaeckman/hermod-api/domain/product"
)

type Controller struct {
	Interactor Interactor
}

type Interactor struct {
	product.Repository
	product.Presenter
}

type ReqBody struct {
	ID    string
	Title string
}

func (i *Interactor) Create(rb ReqBody) error {
	fmt.Println("######[Interactor]########")
	fmt.Println(rb)
	fmt.Println("#########################")

	p := rb.mapEntity()

	p, _ = i.Store(p)
	i.PresentCreated(&p)

	return nil
}

func (i *Interactor) Get(id string) (domain.Product, error) {
	p, _ := i.Get(id)

	i.Present(&p)

	return p, nil
}

func (i *Interactor) GetAll() ([]*domain.Product, error) {
	p, _ := i.GetAll()

	i.PresentAll(p)

	return p, nil
}
