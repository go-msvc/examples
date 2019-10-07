package main

import (
	"github.com/go-msvc/domain"
	"github.com/jansemmelink/log"
)

//oper0 does nothing. This is the minimum you have to code to make a micro-service
type oper0 struct{}

func (r oper0) Exec() domain.Result {
	log.Debugf("oper0.Exec() called...")
	return domain.Success(nil, nil)
}
