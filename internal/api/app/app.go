package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/expx-tech/go-framework/internal/api/config"
	"github.com/expx-tech/go-framework/internal/api/handlers/middleware"
	logClient "github.com/expx-tech/go-framework/pkg/logger"
)

type logger interface {
	Info(msg string)
	Error(err error)
}

func Start() {
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	conf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	log, err := logClient.NewZapLog(conf.Api.LogLevel)
	if err != nil {
		panic(err)
	}

	router := newRouter(log)

	server := &http.Server{
		Addr: ":" + conf.Api.HTTPPort,
		Handler: middleware.AuthHandler(
			middleware.LogHandler(
				router.Get(),
				log,
			),
			log,
		),
	}

	go func() {
		<-ctx.Done()
		if err := server.Close(); err != nil {
			log.Error(fmt.Errorf("%v: %w", errServerShutdown, err))
		}
		log.Info("server shutdown")
	}()

	if err := server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Error(fmt.Errorf("%v: %w", errStartupServer, err))
		}
	}
}
