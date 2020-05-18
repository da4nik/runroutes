package main

import (
	"github.com/da4nik/runroutes/internal/config"
	"github.com/da4nik/runroutes/internal/log"
	"github.com/da4nik/runroutes/internal/rest"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf, err := config.Load("runroutes.yml")
	if err != nil {
		panic(err)
	}

	logger := log.NewLogger("runroutes.log")

	rst := rest.NewServer(conf.Port, logger)

	if err := rst.Start(); err != nil {
		logger.Errorf("Failed to start REST service: %s", err.Error())
		panic(err)
	}
	defer rst.Stop()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch
}
