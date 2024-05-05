package healthcheck

import (
	"net/http"

	"github.com/activehigh/go-gin-project-template/internal/handlers"
)

// healthcheck request
// =====================================================================
type HealthcheckRequest struct {
	*handlers.RawRequest
}

// healthcheck response
// =====================================================================
type HealthcheckResponse struct {
	Message string `json:"message"`
}

// healthcheck handler
// =====================================================================
type HealthcheckHandler struct {
}

func (p HealthcheckHandler) CreateRequestObject() HealthcheckRequest {
	rq := HealthcheckRequest{
		RawRequest: &handlers.RawRequest{},
	}
	return rq
}

func (p HealthcheckHandler) Handle(r HealthcheckRequest) (HealthcheckResponse, int, error) {
	response := HealthcheckResponse{
		Message: "I am alive!",
	}
	return response, http.StatusOK, nil
}
