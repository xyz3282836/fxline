package route

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"go.uber.org/zap"
)

// HelloHandler is an HTTP handler that
// prints a greeting to the user.
type HelloHandler struct {
	log    *zap.Logger
	reqNum int
	lock   sync.Mutex
}

// NewHelloHandler builds a new HelloHandler.
func NewHelloHandler(log *zap.Logger) *HelloHandler {
	return &HelloHandler{log: log}
}

func (h *HelloHandler) Pattern() string {

	return "/hello"
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.lock.Lock()
	var reqNum int
	h.reqNum++
	reqNum = h.reqNum
	h.lock.Unlock()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("Failed to read request", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if _, err := fmt.Fprintf(w, "Hello, body is %s reqNum is %d\n", body, reqNum); err != nil {
		h.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
