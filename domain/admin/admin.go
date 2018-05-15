package admin

import (
	"github.com/RobinBaeckman/hermod-api/domain"
)

// UserRepository is the repo for a product
type Repository interface {
	Get(email string) (domain.Admin, error)
	Persist(domain.Admin) error
}
