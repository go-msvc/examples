package main

import (
	"fmt"

	"github.com/go-msvc/domain"
	"github.com/jansemmelink/log"
)

type greetRequest struct {
	Name    string   `json:"name"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

func (g greetRequest) Validate() error {
	if len(g.Name) == 0 {
		return log.Wrapf(nil, "missing name")
	}
	return nil
}

type greetResponse struct {
	Message string `json:"message"`
}

func (req greetRequest) Exec() domain.Result {
	return domain.Success(
		greetResponse{
			Message: fmt.Sprintf("Hi %s! %d greetings from %v!", req.Name, req.Count, req.Friends),
		},
		nil)
}
