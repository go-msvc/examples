package main

import (
	"github.com/go-msvc/domain"
	"github.com/go-msvc/domain/server/rest"
	"github.com/jansemmelink/log"
)

func main() {
	log.DebugOn()
	domain.New("hello").
		WithOper("greet", greet{}).
		WithOper("sing", sing{}).
		WithConfig("rest", rest.Config{Address: "localhost:12345"}).
		Run()
}

//Greet ...
type greet struct {
	Name string `json:"name"`
}

type greetResponse struct {
	Message string `json:"message"`
}

func (g greet) Validate() error {
	if len(g.Name) == 0 {
		return log.Wrapf(nil, "missing name")
	}
	return nil
}

func (g greet) Run() (domain.IResponse, error) {
	return greetResponse{Message: "Hi " + g.Name + "!"}, nil
}

//Sing ...
type sing struct{}

func (s sing) Validate() error {
	return nil
}

func (s sing) Run() (domain.IResponse, error) {
	return "La-la-la", nil
}
