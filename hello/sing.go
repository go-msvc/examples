package main

import (
	"github.com/go-msvc/domain"
)

//Sing demonstrates the a simple micro-service
//no request params and constant response
type singRequest struct{}

func (s singRequest) Exec(domain.Context, domain.IRequest) domain.IOutcome {
	return domain.NewOutcome(
		nil,        //result code
		"la-la-la", //response
		nil)        //audit
}
