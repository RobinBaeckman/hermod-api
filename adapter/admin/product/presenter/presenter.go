package presenter

import (
	"fmt"

	"github.com/RobinBaeckman/hermod-api/domain"
	"github.com/RobinBaeckman/hermod-api/domain/product"
)

type Presenter struct {
	product.Viewer
}

func (pr Presenter) PresentCreated(p *domain.Product) {
	fmt.Println("######[Presenter]########")
	fmt.Println(p)
	fmt.Println("#########################")
	pr.ViewCreated(p)
}

func (pr Presenter) Present(p *domain.Product) {
	fmt.Println("######[Presenter]########")
	fmt.Println(p)
	fmt.Println("#########################")
	pr.View(p)
}

func (pr Presenter) PresentAll(p []*domain.Product) {
	fmt.Println("######[Presenter]########")
	fmt.Println(p)
	fmt.Println("#########################")
	pr.ViewAll(p)
}
