package main

import (
	"github.com/go-msvc/domain"
)

//Sing demonstrates the a simple micro-service
//no request params and constant response
type singRequest struct{}

func (s singRequest) Exec() domain.Result {
	return domain.Success("La-la-la", nil)
}
