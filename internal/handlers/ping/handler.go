package ping

import (
	"net/http"

	"github.com/activehigh/go-gin-project-template/internal/handlers"
)

// ping request
// =====================================================================
type PingRequest struct {
	*handlers.RawRequest
}

// pong response
// =====================================================================
type PongResponse struct {
	Message string `json:"message"`
}

// ping handler
// =====================================================================
type PingHandler struct {
}

func (p PingHandler) CreateRequestObject() PingRequest {
	rq := PingRequest{
		RawRequest: &handlers.RawRequest{},
	}
	return rq
}

func (p PingHandler) Handle(r PingRequest) (PongResponse, int, error) {
	response := PongResponse{
		Message: "pong",
	}
	return response, http.StatusOK, nil
}
