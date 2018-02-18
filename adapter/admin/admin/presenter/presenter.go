package presenter

import (
	"github.com/RobinBaeckman/hermod-api/usecase/admin/admin"
)

type Presenter struct {
	Viewer
}

type Viewer interface {
	ViewCreated(*ViewModel) (err error)
	View(*ViewModel) (err error)
	ViewAll(*[]ViewModel) (err error)
}

type ViewModel struct {
	ID        string
	Email     string
	Password  string
	FirstName string
	LastName  string
}

func (p Presenter) PresentCreated(outd *admin.OutputData) (err error) {
	vm := &ViewModel{}
	if err = vm.mapper(outd); err != nil {
		return err
	}
	if err = p.ViewCreated(vm); err != nil {
		return err
	}

	return
}

func (p Presenter) Present(outd *admin.OutputData) (err error) {
	vm := &ViewModel{}
	if err = vm.mapper(outd); err != nil {
		return err
	}
	if err = p.View(vm); err != nil {
		return err
	}

	return
}

func (p Presenter) PresentAll(outds *[]admin.OutputData) (err error) {
	vms := []ViewModel{}
	for _, outd := range *outds {
		vm := &ViewModel{}
		if err = vm.mapper(&outd); err != nil {
			return err
		}
		vms = append(vms, *vm)
	}

	if err = p.ViewAll(&vms); err != nil {
		return err
	}

	return
}

func (vm *ViewModel) mapper(outd *admin.OutputData) error {
	*vm = ViewModel{
		ID:        outd.ID,
		Email:     outd.Email,
		Password:  "*****",
		FirstName: outd.FirstName,
		LastName:  outd.LastName,
	}

	return nil
}
