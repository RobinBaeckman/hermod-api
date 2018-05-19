package admin

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
	Persist(*domain.Admin) error
	Get(*domain.Admin) error
	GetAll(*[]domain.Admin) error
}

type Presenter interface {
	PresentStored(*OutputData) (err error)
	Present(*OutputData) (err error)
	PresentAll(*[]OutputData) (err error)
}

func (i *Interactor) Index() (err error) {
	as := &[]domain.Admin{}
	err = i.GetAll(as)
	if err != nil {
		return err
	}
	outds := []OutputData{}
	for _, a := range *as {
		outd := &OutputData{}
		outd.Map(&a)
		outds = append(outds, *outd)
	}
	if err = i.PresentAll(&outds); err != nil {
		return err
	}

	return
}

func (i *Interactor) Show(ind InputDataShow) (err error) {
	a := &domain.Admin{}
	ind.Map(a)
	err = i.Get(a)
	if err != nil {
		return err
	}
	outd := &OutputData{}
	outd.Map(a)
	if err := i.Present(outd); err != nil {
		return err
	}

	return
}

func (i *Interactor) Store(ind InputDataStore) (err error) {
	a := &domain.Admin{}
	if err = ind.Map(a); err != nil {
		return err
	}
	err = i.Persist(a)
	if err != nil {
		return err
	}
	outd := &OutputData{}
	outd.Map(a)
	if err = i.PresentStored(outd); err != nil {
		return err
	}

	return
}
