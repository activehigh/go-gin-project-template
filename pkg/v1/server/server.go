package server

import (
	"context"
	"fmt"
	"github.com/activehigh/go-gin-project-template/pkg/v1/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	server "github.com/activehigh/go-gin-project-template/internal/server"
	"github.com/activehigh/go-gin-project-template/pkg/v1/configs"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

// Server represents the main server instance
type Server struct {
	config *configs.Config
}

// NewServer creates a new server instance
func NewServer(config *configs.Config) *Server {
	return &Server{
		config: config,
	}
}

// Start initializes and starts the server
func (s *Server) Start() error {
	logger.InitializeLogger()
	log := zap.L()

	log.Debug(fmt.Sprintf("config loaded: %+v", s.config))
	router := gin.New()

	// get router with routes setup
	router = server.SetHealthchecks(s.config, router)

	// configure logger for router
	router.Use(ginzap.Ginzap(log, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(log, true))

	// set all other routes
	router = server.SetRoutes(s.config, router)

	// start server
	cw := make([]server.ConnectionWatcher, 2)
	cw = append(cw, server.ConnectionWatcher{})
	srv := &http.Server{
		Addr:      fmt.Sprintf(":%d", s.config.Port),
		Handler:   router,
		ConnState: cw[0].OnStateChange,
	}
	servers := [](*http.Server){srv}

	// setup graceful termination of server
	for _, s := range servers {
		go func(srv *http.Server) {
			// service connections
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatal(fmt.Sprintf("listen: %s\n", err))
			}
		}(s)
	}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down Server ...")

	terminationPeriod := time.Duration(s.config.TerminationGracePeriodInSeconds)
	ctx, cancel := context.WithTimeout(
		context.Background(), terminationPeriod*time.Second,
	)
	defer cancel()

	// shutdown servers, which should be pretty easy as no active conenction should be there
	for i, s := range servers {
		go func(index int, srv *http.Server) {
			// wait to give time to proxy to finish first
			log.Info(
				fmt.Sprintf(
					"Active connection for server: %v is: %v", index,
					cw[index].Count(),
				),
			)
			if err := srv.Shutdown(ctx); err != nil {
				log.Info(fmt.Sprintf("Error on server shutdown: %s", err))
			} else {
				log.Info("Server shutdown succeeded")
			}
		}(i, s)
	}

	<-ctx.Done()
	log.Info("Server exiting")
	return nil
}
