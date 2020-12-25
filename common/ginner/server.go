package ginner

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Linus-Boehm/go-serverless-suite/common"

	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog"
)

type HTTPServer interface {
	ListenAndServe() error
	Shutdown(context.Context) error
	SetKeepAlivesEnabled(bool)
}

type Servers struct {
	servers []Server
	quit    chan os.Signal
}

func NewServers(servers ...Server) Servers {
	srvs := Servers{
		servers: servers,
		quit:    make(chan os.Signal, 1),
	}
	return srvs
}

func (srvs Servers) StartAllAndWait() {
	for _, server := range srvs.servers {
		func() {
			server.Start()
		}()
	}

	signal.Notify(srvs.quit, syscall.SIGINT, syscall.SIGTERM)
	<-srvs.quit

	for _, server := range srvs.servers {
		server.Shutdown()
	}
}

type Server struct {
	name   string
	server HTTPServer
	port   string
}

// NewServer initializes a new management server on the given port.
// The management server serves endpoints for health and readiness endpoints.
func NewServer(name string, router *gin.Engine, port string, errorLogger zerolog.Logger) Server {
	return Server{
		name: name,
		server: &http.Server{
			Handler:      router,
			Addr:         port,
			ErrorLog:     log.New(errorLogger, "", int(zerolog.ErrorLevel)),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
		port: port,
	}
}

// Run will start the management server asynchronous and returns immediately.
func (srv *Server) Start() {
	go func(server HTTPServer) {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srv.logError("unable to start server", err)
		}
	}(srv.server)

	srv.logInfo("server is ready to handle requests")
}

func (srv *Server) Shutdown() {
	srv.logInfo("server is shutting down")

	// use timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.server.SetKeepAlivesEnabled(false)
	if err := srv.server.Shutdown(ctx); err != nil {
		srv.logError("could not gracefully shutdown the server", err)
	} else {
		srv.logInfo("server stopped")
	}
}

func (srv *Server) logInfo(message string) {
	common.GetDefaultLogger().Info().Msg(message)
}

func (srv *Server) logError(message string, err error) {
	common.GetDefaultLogger().Error().Err(err).Msg(message)
}
