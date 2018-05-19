package auth

import (
	"encoding/json"

	"github.com/RobinBaeckman/hermod-api/domain"
	"github.com/pkg/errors"
)

type InputDataAuth struct {
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

func (outd OutputData) Map(a *domain.Admin) error {
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
