package main

import (
	"github.com/go-msvc/domain"
	"github.com/go-msvc/domain/server/rest"
)

func main() {
	domain.New("hello").
		WithOper("greet", greet{}).
		WithOper("sing", sing{}).
		WithConfig("rest", rest.Config{Address: "localhost:12345"}).
		Run()
}

//Greet ...
type greet struct{}

func (g greet) Run() (domain.IResponse, error) {
	return "Hi", nil
}

//Sing ...
type sing struct{}

func (s sing) Run() (domain.IResponse, error) {
	return "La-la-la", nil
}
