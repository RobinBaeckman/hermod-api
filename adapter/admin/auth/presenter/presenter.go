package presenter

import (
	"encoding/json"

	"github.com/RobinBaeckman/hermod-api/usecase/admin/auth"
	"github.com/pkg/errors"
)

type Presenter struct {
	Viewer
}

type Viewer interface {
	View(ViewModel) error
}

type ViewModel struct {
	ID        string
	Email     string
	Password  []byte
	FirstName string
	LastName  string
}

func (p Presenter) Present(outd auth.OutputData) (err error) {
	vm := ViewModel{}

	if err := vm.mapper(&outd); err != nil {
		return err
	}

	if err = p.View(vm); err != nil {
		return err
	}

	return err
}

func (vm ViewModel) mapper(outd *auth.OutputData) (err error) {
	bs, err := json.Marshal(&outd)
	if err != nil {
		return errors.Wrap(err, "Failed marshaling admin to bytestring")
	}
	err = json.Unmarshal(bs, &vm)
	if err != nil {
		return errors.Wrap(err, "Failed unmarshaling bytestring to outputdata")
	}

	return err
}
