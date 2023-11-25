package main

import (
	"fmt"
	"github.com/aakash-rajur/fold-se/internal/env"
	"github.com/aakash-rajur/fold-se/internal/es"
	"github.com/aakash-rajur/fold-se/internal/routes"
	"github.com/aakash-rajur/fold-se/internal/store"
	_ "github.com/lib/pq"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	workdir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	envVars := env.Load(workdir)

	db, dbClose, err := store.Connect(envVars)

	if err != nil {
		panic(err)
	}

	defer dbClose()

	esc, err := es.Client(envVars)

	server := routes.Router(routes.Args{Db: db, Esc: esc})

	address := fmt.Sprintf("0.0.0.0:%s", envVars["PORT"])

	serveErrChan := make(chan error)

	serve := func(errChan chan error, address string) {
		errChan <- server.Run(address)
	}

	go serve(serveErrChan, address)

	interrupt := make(chan os.Signal)

	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-serveErrChan:
		{
			panic(err)

			return
		}
	case <-interrupt:
		return
	}
}
