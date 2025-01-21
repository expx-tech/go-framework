package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type logger interface {
	Error(err error)
}

type Health struct {
	log logger
}

func NewHealth(opts ...func(*Health)) *Health {
	h := new(Health)
	for _, opt := range opts {
		opt(h)
	}

	return h
}

func WithHealthLogger(l logger) func(*Health) {
	return func(h *Health) {
		h.log = l
	}
}

func (h *Health) Healthz(c *gin.Context) {
	c.String(http.StatusOK, "")
}
