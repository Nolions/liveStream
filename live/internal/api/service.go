package api

import (
	"context"
	"errors"
	"fmt"
	"live/conf"
	log "live/internal/helper"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type App struct {
	Ctx context.Context
}

type Server struct {
	HttpServer *http.Server
}

func New(ctx context.Context, conf *conf.App) *Server {
	app := App{
		Ctx: ctx,
	}

	e := engine(conf.Debug)
	app.router(e)
	addr := fmt.Sprintf(":%s", conf.Addr)
	h := &http.Server{
		Addr:         addr,
		Handler:      e,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return &Server{
		HttpServer: h,
	}
}

func (serv *Server) Run() {
	log.Info("Starting HTTP server", zap.String("address", serv.HttpServer.Addr))
	if err := serv.HttpServer.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		log.Error("Failed to start HTTP server", zap.Error(err))
	} else {
		log.Info("HTTP server stopped")
	}
}

func (serv *Server) Shutdown(ctx context.Context) error {
	log.Info("Shutting down HTTP server")
	err := serv.HttpServer.Shutdown(ctx)
	if err != nil {
		log.Error("HTTP server shutdown error", zap.Error(err))
		return err
	}

	log.Info("HTTP server shutdown successfully")
	return nil
}
