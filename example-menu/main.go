package main

import (
	"fmt"
	"net/http"

	"github.com/go-msvc/log"
	"github.com/go-msvc/service"
	svchttp "github.com/go-msvc/service/http"
)

func main() {
	//switch logging on for all loggers
	log.Top().SetLevel(log.DebugLevel)

	//define a simple menu:
	menu := newMenu("Main").
		with("Hello").
		with("Profile").
		with("Exit")

	//create a service to serve that menu
	svc := service.New().WithOper("menu", menu)
	if err := http.ListenAndServe("localhost:12345", svchttp.Handler(svc)); err != nil {
		panic(err)
	}
}

func newMenu(title string) menu {
	return menu{
		title: title,
		items: []menuItem{},
	}
}

type menu struct {
	title string
	items []menuItem
}

type menuItem struct {
	label string
}

func (m menu) with(label string) menu {
	m.items = append(m.items, menuItem{label: label})
	return m
}

//Handle implement service.IOper.Handle()
func (m menu) Handle(ctx service.IContext, req interface{}) (service.IResponse, error) {
	text := m.title
	for index, i := range m.items {
		text += fmt.Sprintf("\n%d) %s", index+1, i.label)
	}
	return text, nil
}
