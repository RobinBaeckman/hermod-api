package product

import "github.com/RobinBaeckman/hermod-api/domain"

// ProductRepository is the repo for a product
type Repository interface {
	Store(p domain.Product) (domain.Product, error)
	Get(id string) (domain.Product, error)
	GetAll() ([]*domain.Product, error)
}

type Presenter interface {
	PresentCreated(*domain.Product)
	Present(*domain.Product)
	PresentAll([]*domain.Product)
}

type Viewer interface {
	ViewCreated(*domain.Product)
	View(*domain.Product)
	ViewAll([]*domain.Product)
}
