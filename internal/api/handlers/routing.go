package handlers

import (
	"github.com/gin-gonic/gin"
)

type healthHandler interface {
	Healthz(c *gin.Context)
}

type Router struct {
	client *gin.Engine
	health healthHandler
}

func NewRouter(opts ...func(*Router)) *Router {
	r := new(Router)
	r.client = gin.New()

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func WithHealthHandler(hh healthHandler) func(*Router) {
	return func(r *Router) {
		r.health = hh
	}
}

func (r *Router) Get() *gin.Engine {
	r.client.GET("/healthz", r.health.Healthz)

	return r.client
}
