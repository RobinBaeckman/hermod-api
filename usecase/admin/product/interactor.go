package product

import (
	"github.com/RobinBaeckman/hermod-api/domain"
)

type Controller struct {
	Interactor Interactor
}

type Interactor struct {
	Repository
	Presenter
}

type Repository interface {
	Persist(*domain.Product) error
	Get(*domain.Product) error
	GetAll(*[]domain.Product) error
}

type Presenter interface {
	PresentStored(*OutputData) (err error)
	Present(*OutputData) (err error)
	PresentAll(*[]OutputData) (err error)
}

func (i *Interactor) Index() (err error) {
	ps := &[]domain.Product{}
	err = i.GetAll(ps)
	if err != nil {
		return err
	}
	outds := []OutputData{}
	for _, p := range *ps {
		outd := &OutputData{}
		outd.Map(&p)
		outds = append(outds, *outd)
	}
	if err = i.PresentAll(&outds); err != nil {
		return err
	}

	return
}

func (i *Interactor) Show(ind InputDataShow) (err error) {
	p := &domain.Product{}
	ind.Map(p)
	err = i.Get(p)
	if err != nil {
		return err
	}
	outd := &OutputData{}
	outd.Map(p)
	if err := i.Present(outd); err != nil {
		return err
	}

	return
}

func (i *Interactor) Store(ind InputDataStore) (err error) {
	p := &domain.Product{}
	if err = ind.Map(p); err != nil {
		return err
	}
	err = i.Persist(p)
	if err != nil {
		return err
	}
	outd := &OutputData{}
	outd.Map(p)
	if err = i.PresentStored(outd); err != nil {
		return err
	}

	return
}
