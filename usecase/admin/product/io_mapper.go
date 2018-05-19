package product

import (
	"github.com/RobinBaeckman/hermod-api/domain"
)

type InputDataStore struct {
	Title string
}

type InputDataShow struct {
	ID string
}

type OutputData struct {
	ID    string
	Title string
}

func (outd *OutputData) Map(p *domain.Product) {
	*outd = OutputData{
		ID:    p.ID,
		Title: p.Title,
	}
}

func (ind InputDataStore) Map(a *domain.Product) error {
	*a = domain.Product{
		Title: ind.Title,
	}

	return nil
}

func (ind InputDataShow) Map(a *domain.Product) {
	*a = domain.Product{
		ID: ind.ID,
	}
}
