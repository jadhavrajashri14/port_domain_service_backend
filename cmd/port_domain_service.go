package main

import (
	"context"
	"os"
	"os/signal"
	"port_domain_service_backend/internal/adapters/handler"
	"port_domain_service_backend/internal/adapters/repository"
	"port_domain_service_backend/internal/core/services"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	sigdone := make(chan os.Signal, 1)
	signal.Notify(sigdone, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	portDomains := make(map[string]map[string]interface{})
	repo := repository.PortDomainsRepository{
		PortDomains: portDomains,
	}
	pds := services.NewPortDomainsService(repo)

	svr, err := handler.NewServer(ctx, cancel, pds)
	if err != nil {
		logrus.WithError(err).Fatal("unable to create the API server")
	}

	go func() {
		svr.Run()
		cancel()
	}()

	select {
	case <-ctx.Done():
		logrus.WithError(ctx.Err()).Info("main got cancel")
	case <-sigdone:
		logrus.Info("got sigdone, sending cancel")
		cancel()
	}

	logrus.Info("exiting port_domain_service_backend")
}
