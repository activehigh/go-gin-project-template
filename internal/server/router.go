package server

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"strings"

	"github.com/activehigh/go-gin-project-template/configs"
	handlers "github.com/activehigh/go-gin-project-template/internal/handlers"
	"github.com/gin-gonic/gin"

	"github.com/activehigh/go-gin-project-template/internal/handlers/healthcheck"
	ping "github.com/activehigh/go-gin-project-template/internal/handlers/ping"
)

// method to check if we should ignore path
func ShouldIgnoreRoute(path string) bool {
	routesToIgnore := []string{"/live", "/ready"}
	for _, r := range routesToIgnore {
		if strings.HasPrefix(path, r) {
			return true
		}
	}
	return false
}

// wrapper to resolve generic request, response and handler types
func BindHandler[TReq handlers.Request, TRes handlers.Response](handler handlers.Handler[TReq, TRes]) func(c *gin.Context) {
	logger := zap.L()

	defer func() {
		if err := recover(); err != nil {
			logger.Error(fmt.Sprintf("Unexpected panic error: %v", err))
		}
	}()

	handlerWrapperFunc := func(c *gin.Context) {
		_ = ShouldIgnoreRoute(c.Request.URL.Path)
		t := handler.CreateRequestObject()
		if err := t.Bind(c); err == nil {
			if response, status, err := handler.Handle(t); err == nil {
				c.JSON(status, response)
			} else {
				c.JSON(status, map[string]string{"error": err.Error()})
			}
		} else {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
	return handlerWrapperFunc

}

func SetHealthchecks(config configs.CliConfig, router *gin.Engine) *gin.Engine {
	// readiness & liveness probes
	router.GET(
		"/ready",
		BindHandler[
			healthcheck.HealthcheckRequest,
			healthcheck.HealthcheckResponse,
		](healthcheck.HealthcheckHandler{}),
	)

	router.GET(
		"/live",
		BindHandler[healthcheck.HealthcheckRequest, healthcheck.HealthcheckResponse](healthcheck.HealthcheckHandler{}),
	)

	return router
}

func SetRoutes(c configs.CliConfig, router *gin.Engine) *gin.Engine {
	// ping
	router.GET("/ping", BindHandler[ping.PingRequest, ping.PongResponse](ping.PingHandler{}))
	router.POST("/ping", BindHandler[ping.PingRequest, ping.PongResponse](ping.PingHandler{}))

	return router
}
