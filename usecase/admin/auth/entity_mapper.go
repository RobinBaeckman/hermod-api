package auth

import (
	"encoding/json"
	"fmt"

	"github.com/RobinBaeckman/hermod-api/domain"
)

func (ind InputData) mapEntity() domain.Admin {
	bs, err := json.Marshal(ind)
	if err != nil {
		fmt.Println(err)
	}

	a := domain.Admin{}
	err = json.Unmarshal(bs, &a)
	if err != nil {
		fmt.Println(err)
	}

	return a
}
