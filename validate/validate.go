package validate

import (
	"reflect"

	"github.com/RobinBaeckman/hermod-api/customerr"
)

func Check(ind interface{}) error {
	v := reflect.ValueOf(ind)
	if reflect.ValueOf(ind).Kind() == reflect.Ptr {
		v = reflect.ValueOf(ind).Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == "" {
			return &customerr.App{nil, "Missing parameters", 404}
		}
	}

	return nil
}
