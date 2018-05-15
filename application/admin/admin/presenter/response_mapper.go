package presenter

import "github.com/RobinBaeckman/hermod-api/usecase/admin/admin"

func (vm *ViewModel) Map(outd *admin.OutputData) error {
	*vm = ViewModel{
		ID:        outd.ID,
		Email:     outd.Email,
		Password:  "*****",
		FirstName: outd.FirstName,
		LastName:  outd.LastName,
	}

	return nil
}
