package route

import (
	"io"
	"net/http"

	"go.uber.org/zap"
)

type Route interface {
	http.Handler

	// Pattern reports the path at which this is registered.
	Pattern() string
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}

// EchoHandler is an http.Handler that copies its request body
// back to the response.
type EchoHandler struct {
	log *zap.Logger
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

// ServeHTTP handles an HTTP request to the /echo endpoint.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.log.Warn("Failed to handle request", zap.Error(err))
	}
}
