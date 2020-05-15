package main

import (
	"net/http"

	"github.com/go-msvc/service"
	svchttp "github.com/go-msvc/service/http"
)

func main() {
	svc := service.New().WithOper("hello", hello{})
	if err := http.ListenAndServe("localhost:12345", svchttp.Handler(svc)); err != nil {
		panic(err)
	}
}

//hello implements service.IOper
type hello struct{}

func (h hello) Handle(ctx service.IContext, req interface{}) (service.IResponse, error) {
	return "Hello", nil
}
