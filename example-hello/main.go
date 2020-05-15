package main

import (
	"github.com/go-msvc/service"

	//include at least one server implementation
	//http is used by default on localhost:8080
	//without config
	_ "github.com/go-msvc/service/http"
)

func main() {
	//uncomment to see all debug: log.Top().SetLevel(log.DebugLevel)
	svc := service.New().WithOper("hello", hello{})
	if err := svc.Run(); err != nil {
		panic(err)
	}
}

//hello implements service.IOper
type hello struct{}

func (h hello) Handle(ctx service.IContext, req interface{}) (service.IResponse, error) {
	return "Hello", nil
}
