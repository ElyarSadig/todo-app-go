package server

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/elyarsadig/todo-app/pkg/logger"
	"github.com/nahojer/httprouter"
)

const (
	address        = ":8080"
	writeTimeout   = time.Second * 1
	readTimeout    = time.Second * 1
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

type Server struct {
	logger logger.Logger
	router *httprouter.Router
	db     *sql.DB
}

func New(db *sql.DB, logger logger.Logger) *Server {
	router := httprouter.New().Group("apis/v1")
	router.ErrorHandler = errorHandler
	return &Server{
		db:     db,
		logger: logger,
		router: router,
	}
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:           address,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		s.logger.Info(fmt.Sprintf("Server is listening on PORT: %s", srv.Addr))
		if err := srv.ListenAndServe(); err != nil {
			s.logger.Fatal(err.Error())
		}
	}()

	//TODO map handlers here

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout)
	defer shutdown()

	s.logger.Info("Server Exited Properly")

	return srv.Shutdown(ctx)
}
