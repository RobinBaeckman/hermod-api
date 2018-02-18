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
	Store(*domain.Admin) error
	GetSingle(*domain.Admin) error
	GetMany(*[]domain.Admin) error
}

type Presenter interface {
	PresentCreated(*OutputData) (err error)
	Present(*OutputData) (err error)
	PresentAll(*[]OutputData) (err error)
}

func (i *Interactor) GetAll() (err error) {
	as := &[]domain.Admin{}
	err = i.GetMany(as)
	if err != nil {
		return err
	}
	outds := []OutputData{}
	for _, a := range *as {
		outd := &OutputData{}
		outd.mapper(&a)
		outds = append(outds, *outd)
	}
	if err = i.PresentAll(&outds); err != nil {
		return err
	}

	return
}

func (i *Interactor) Get(ind InputData) (err error) {
	a := &domain.Admin{}
	ind.mapper(a)
	err = i.GetSingle(a)
	if err != nil {
		return err
	}
	outd := &OutputData{}
	outd.mapper(a)
	if err := i.Present(outd); err != nil {
		return err
	}

	return
}

func (i *Interactor) Create(ind InputData) (err error) {
	a := &domain.Admin{}
	if err = ind.createMapper(a); err != nil {
		return err
	}
	outd := &OutputData{}
	err = i.Store(a)
	if err != nil {
		return err
	}
	outd.mapper(a)
	if err = i.Present(outd); err != nil {
		return err
	}

	return
}
