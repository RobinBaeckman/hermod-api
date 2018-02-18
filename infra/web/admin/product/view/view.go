package view

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RobinBaeckman/hermod-api/domain"
)

type Viewer struct {
	http.ResponseWriter
}

func (v Viewer) ViewCreated(p *domain.Product) {
	fmt.Println("######[View]########")
	fmt.Println(p)
	fmt.Println("#########################")
	jbs, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	v.Header().Set("Content-Type", "application/json")
	v.WriteHeader(http.StatusOK)
	v.Write(jbs)
}

func (v Viewer) View(p *domain.Product) {
	fmt.Println("######[View]########")
	fmt.Println(p)
	fmt.Println("#########################")
	jbs, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	v.Header().Set("Content-Type", "application/json")
	v.WriteHeader(http.StatusOK)
	v.Write(jbs)
}

func (v Viewer) ViewAll(p []*domain.Product) {
	fmt.Println("######[View]########")
	fmt.Println(p)
	fmt.Println("#########################")
	jbs, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	v.Header().Set("Content-Type", "application/json")
	v.WriteHeader(http.StatusOK)
	v.Write(jbs)
}
