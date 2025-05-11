package server

import (
	"encoding/json"
	"github.com/activehigh/go-gin-project-template/pkg/v1/configs"
	"net/http"
	"net/http/httptest"
	"testing"

	healthcheck "github.com/activehigh/go-gin-project-template/internal/handlers/healthcheck"
	ping "github.com/activehigh/go-gin-project-template/internal/handlers/ping"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRoutes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Routes Test Suite")
}

var _ = Describe(
	"Route", func() {
		// test route - /ping
		It(
			"/ping will return pong", func() {
				router := gin.New()
				router = SetHealthchecks(&configs.Config{}, router)
				router = SetRoutes(&configs.Config{}, router)

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/ping", nil)
				router.ServeHTTP(w, req)

				response := ping.PongResponse{}
				_ = json.Unmarshal(w.Body.Bytes(), &response)

				Expect(w.Code).To(Equal(200))
				Expect(response.Message).To(Equal("pong"))
			},
		)

		// test route - /live
		It(
			"/live will return I am alive", func() {
				router := gin.New()
				router = SetHealthchecks(&configs.Config{}, router)
				router = SetRoutes(&configs.Config{}, router)

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/live", nil)
				router.ServeHTTP(w, req)

				response := healthcheck.HealthcheckResponse{}
				_ = json.Unmarshal(w.Body.Bytes(), &response)

				Expect(w.Code).To(Equal(200))
				Expect(response.Message).To(Equal("I am alive!"))
			},
		)

		// test route - /ready
		It(
			"/ready will return I am alive", func() {
				router := gin.New()
				router = SetHealthchecks(&configs.Config{}, router)
				router = SetRoutes(&configs.Config{}, router)

				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/ready", nil)
				router.ServeHTTP(w, req)

				response := healthcheck.HealthcheckResponse{}
				_ = json.Unmarshal(w.Body.Bytes(), &response)

				Expect(w.Code).To(Equal(200))
				Expect(response.Message).To(Equal("I am alive!"))
			},
		)
	},
)
