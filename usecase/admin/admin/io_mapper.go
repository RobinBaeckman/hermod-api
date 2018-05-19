package admin

import (
	"github.com/RobinBaeckman/hermod-api/domain"
	"golang.org/x/crypto/bcrypt"
)

type InputDataStore struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type InputDataShow struct {
	ID string
}

type OutputData struct {
	ID        string
	Email     string
	Password  string
	FirstName string
	LastName  string
}

func (outd *OutputData) Map(a *domain.Admin) {
	*outd = OutputData{
		ID:        a.ID,
		Email:     a.Email,
		Password:  "*****",
		FirstName: a.FirstName,
		LastName:  a.LastName,
	}
}

func (ind InputDataStore) Map(a *domain.Admin) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(ind.Password), bcrypt.DefaultCost)
	a.Password = hash
	if err != nil {
		return err
	}
	*a = domain.Admin{
		Email:     ind.Email,
		Password:  hash,
		FirstName: ind.FirstName,
		LastName:  ind.LastName,
	}

	return nil
}

func (ind InputDataShow) Map(a *domain.Admin) {
	*a = domain.Admin{
		ID: ind.ID,
	}
}
