package product

import (
	"encoding/json"
	"fmt"

	"github.com/RobinBaeckman/hermod-api/domain"
)

func (rb ReqBody) mapEntity() domain.Product {
	bs, err := json.Marshal(rb)
	if err != nil {
		fmt.Println(err)
	}

	p := domain.Product{}
	err = json.Unmarshal(bs, &p)
	if err != nil {
		fmt.Println(err)
	}

	return p
}
