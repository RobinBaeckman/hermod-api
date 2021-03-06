package presenter

import "github.com/RobinBaeckman/hermod-api/usecase/admin/product"

type Presenter struct {
	Viewer
}

type Viewer interface {
	ViewStored(*ViewModel) (err error)
	View(*ViewModel) (err error)
	ViewAll(*[]ViewModel) (err error)
}

type ViewModel struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (p Presenter) PresentStored(outd *product.OutputData) (err error) {
	vm := &ViewModel{}
	if err = vm.Map(outd); err != nil {
		return err
	}
	if err = p.ViewStored(vm); err != nil {
		return err
	}

	return
}

func (p Presenter) Present(outd *product.OutputData) (err error) {
	vm := &ViewModel{}
	if err = vm.Map(outd); err != nil {
		return err
	}
	if err = p.View(vm); err != nil {
		return err
	}

	return
}

func (p Presenter) PresentAll(outds *[]product.OutputData) (err error) {
	vms := []ViewModel{}
	for _, outd := range *outds {
		vm := &ViewModel{}
		if err = vm.Map(&outd); err != nil {
			return err
		}
		vms = append(vms, *vm)
	}

	if err = p.ViewAll(&vms); err != nil {
		return err
	}

	return
}
