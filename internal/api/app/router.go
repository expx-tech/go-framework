package app

import (
	"github.com/expx-tech/go-framework/internal/api/handlers"
	"github.com/expx-tech/go-framework/internal/api/handlers/health"
)

func newRouter(l logger) *handlers.Router {
	hh := health.NewHealth(
		health.WithHealthLogger(l),
	)
	router := handlers.NewRouter(
		handlers.WithHealthHandler(hh),
	)

	return router
}
