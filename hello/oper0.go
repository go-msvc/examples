package main

import (
	"github.com/go-msvc/domain"
	"github.com/go-msvc/log"
)

//oper0 does nothing. This is the minimum you have to code to make a micro-service
type oper0 struct{}

func (r oper0) Exec(ctx domain.Context, req domain.IRequest) domain.IOutcome {
	log.Debugf("oper0.Exec() called...")
	return domain.NewOutcome(
		nil, //result code
		nil, //response
		nil) //audit
}
