package auth

import (
	"github.com/RobinBaeckman/hermod-api/domain"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	Interactor Interactor
}

type Interactor struct {
	Repository
	Presenter
}

type Repository interface {
	Get(email string) (domain.Admin, error)
	Persist(domain.Admin) error
}

type Presenter interface {
	Present(OutputData) error
}

func (i *Interactor) Auth(ind InputDataAuth) (err error) {
	outd := OutputData{}
	a, err := i.Get(ind.Email)
	if err != nil {
		return err
	}

	if err = bcrypt.CompareHashAndPassword(a.Password, []byte(ind.Password)); err != nil {
		return err
	}

	if err := outd.Map(&a); err != nil {
		return err
	}

	if err = i.Present(outd); err != nil {
		return err
	}

	return nil
}
