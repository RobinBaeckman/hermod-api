package admin

import (
	"github.com/RobinBaeckman/hermod-api/domain"
	"golang.org/x/crypto/bcrypt"
)

func (outd *OutputData) mapper(a *domain.Admin) {
	*outd = OutputData{
		ID:        a.ID,
		Email:     a.Email,
		Password:  "*****",
		FirstName: a.FirstName,
		LastName:  a.LastName,
	}
}

func (ind InputData) createMapper(a *domain.Admin) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(ind.Password), bcrypt.DefaultCost)
	a.Password = hash
	if err != nil {
		return err
	}
	*a = domain.Admin{
		ID:        ind.ID,
		Email:     ind.Email,
		Password:  hash,
		FirstName: ind.FirstName,
		LastName:  ind.LastName,
	}

	return nil
}

func (ind InputData) mapper(a *domain.Admin) {
	*a = domain.Admin{
		ID:    ind.ID,
		Email: ind.Email,
	}
}
