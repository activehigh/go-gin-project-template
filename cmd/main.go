package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/activehigh/go-gin-project-template/configs"
	log "github.com/activehigh/go-gin-project-template/internal/logger"
	server "github.com/activehigh/go-gin-project-template/internal/server"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

var config configs.CliConfig

func init() {
	config = configs.LoadArguments()
}

func main() {
	log.InitializeLogger()
	logger := zap.L()

	logger.Debug(fmt.Sprintf("config loaded: %+v", config))
	router := gin.New()
	// =======================================================================================
	// get router with routes setup
	router = server.SetHealthchecks(config, router)

	// configure logger for router
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))

	// set all other routes
	router = server.SetRoutes(config, router)

	// =======================================================================================
	// start server
	// =======================================================================================
	cw := make([]server.ConnectionWatcher, 2)
	cw = append(cw, server.ConnectionWatcher{})
	srv := &http.Server{
		Addr:      ":8080",
		Handler:   router,
		ConnState: cw[0].OnStateChange,
	}
	servers := [](*http.Server){srv}

	// =======================================================================================
	// setup graceful termination of server
	// =======================================================================================
	for _, s := range servers {
		go func(srv *http.Server) {
			// service connections
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Fatal(fmt.Sprintf("listen: %s\n", err))
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
	logger.Info("Shutting down Server ...")

	terminationPeriod := time.Duration(config.TerminationGracePeriodInSeconds)
	ctx, cancel := context.WithTimeout(
		context.Background(), terminationPeriod*time.Second,
	)
	defer cancel()

	// shutdown servers, which should be pretty easy as no active conenction should be there
	for i, s := range servers {
		go func(index int, srv *http.Server) {
			// wait to give time to proxy to finish first
			logger.Info(
				fmt.Sprintf(
					"Active connection for server: %v is: %v", index,
					cw[index].Count(),
				),
			)
			if err := srv.Shutdown(ctx); err != nil {
				logger.Info(fmt.Sprintf("Error on server shutdown: %s", err))
			} else {
				logger.Info("Server shutdown succeeded")
			}
		}(i, s)
	}

	<-ctx.Done()
	logger.Info("Server exiting")
	// =======================================================================================
}
