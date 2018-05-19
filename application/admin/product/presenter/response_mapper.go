package presenter

import "github.com/RobinBaeckman/hermod-api/usecase/admin/product"

func (vm *ViewModel) Map(outd *product.OutputData) error {
	*vm = ViewModel{
		ID:    outd.ID,
		Title: outd.Title,
	}

	return nil
}
