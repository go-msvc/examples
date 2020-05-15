package main

import (
	"github.com/go-msvc/config"
	"github.com/go-msvc/config/source/files"
	"github.com/go-msvc/domain"

	//_ "github.com/go-msvc/domain/server/nats"
	_ "github.com/go-msvc/domain/server/rest"
	"github.com/go-msvc/log"
)

func main() {
	log := log.NewLogger()

	config := config.NewSources().
		With(files.New("./conf"))
		//default static config used if not found in above config sources
		//With(static.New("hello.server.rest", rest.Config{Address: "localhost:12345"}))

	err := domain.New("hello", 1300).
		WithOper("null", oper0{}).
		WithOper("greet", greetRequest{}).
		WithOper("sing", singRequest{}).
		WithConfig(config).
		Run(log)
	if err != nil {
		log.Errorf("Failed %s", err)
	} else {
		log.Debugf("Terminated.")
	}
} //main()
