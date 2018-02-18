package auth

import (
	"encoding/json"

	"github.com/RobinBaeckman/hermod-api/domain"
	"github.com/RobinBaeckman/hermod-api/domain/admin"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	Interactor Interactor
}

type Interactor struct {
	admin.Repository
	Presenter
}

type Presenter interface {
	Present(OutputData) error
}

type InputData struct {
	Email    string
	Password string
}

type OutputData struct {
	ID        string
	Email     string
	Password  []byte
	FirstName string
	LastName  string
}

func (i *Interactor) Auth(ind InputData) (err error) {
	outd := OutputData{}
	a, err := i.Get(ind.Email)
	if err != nil {
		return err
	}

	if err = bcrypt.CompareHashAndPassword(a.Password, []byte(ind.Password)); err != nil {
		return err
	}

	if err := outd.mapper(&a); err != nil {
		return err
	}

	if err = i.Present(outd); err != nil {
		return err
	}

	return nil
}

func (outd OutputData) mapper(a *domain.Admin) error {
	bs, err := json.Marshal(&a)
	if err != nil {
		return errors.Wrap(err, "Failed marshaling admin to bytestring")
	}
	err = json.Unmarshal(bs, &outd)
	if err != nil {
		return errors.Wrap(err, "Failed unmarshaling bytestring to outputdata")
	}

	return nil
}
