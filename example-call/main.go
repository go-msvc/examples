package main

import (
	"net/http"

	"github.com/go-msvc/log"
	"github.com/go-msvc/service"
	svchttp "github.com/go-msvc/service/http"
)

func main() {
	//switch logging on for all loggers
	log.Top().SetLevel(log.DebugLevel)

	svc := service.New().WithOper("relay", relay{})
	if err := http.ListenAndServe("localhost:12346", svchttp.Handler(svc)); err != nil {
		panic(err)
	}
}

//relay implements service.IOper
type relay struct{}

func (h relay) Handle(ctx service.IContext, req interface{}) (service.IResponse, error) {
	res, err := ctx.Call("hello", req)
	return res, err
}
