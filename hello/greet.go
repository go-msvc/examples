package main

import (
	"fmt"

	"github.com/go-msvc/domain"
	"github.com/go-msvc/errors"
)

type greetRequest struct {
	Name    string   `json:"name"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

func (g greetRequest) Validate() error {
	if len(g.Name) == 0 {
		return errors.Wrapf(nil, "missing name")
	}
	return nil
}

type greetResponse struct {
	Message string `json:"message"`
}

func (g greetRequest) Exec(domain.Context, domain.IRequest) domain.IOutcome {
	return domain.NewOutcome(
		nil, //domain.Success(),
		greetResponse{
			Message: fmt.Sprintf("Hi %s! %d greetings from %v!", g.Name, g.Count, g.Friends),
		},
		nil)
}
