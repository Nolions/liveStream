package api

import (
	"live/internal/exception"
	log "live/internal/helper"
	"live/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app App) router(e *gin.Engine) {
	e.HandleMethodNotAllowed = true
	e.NoMethod(response.HandleNoAllowMethod)
	e.NoRoute(response.HandleNotFound)

	e.GET("/healthz", exception.ErrHandler(app.healthHandler))

	event := e.Group("/event")
	{
		event.POST("on_connect", exception.ErrHandler(app.connectEvent))
		event.POST("on_publish", exception.ErrHandler(app.publishEvent))
		event.POST("on_publish_done", exception.ErrHandler(app.PublishDoneEvent))
		event.POST("on_done", exception.ErrHandler(app.doneEvent))
		event.POST("on_play", exception.ErrHandler(app.playEvent))
		event.POST("on_play_done", exception.ErrHandler(app.playDoneEvent))
		event.POST("on_update", exception.ErrHandler(app.updateEvent))
	}
}

func (app App) healthHandler(c *gin.Context) error {
	log.Info("health check endpoint accessed")
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})

	return nil
}

func (app App) connectEvent(c *gin.Context) error {
	log.Info("connectEvent")
	return nil
}

func (app App) doneEvent(c *gin.Context) error {
	log.Info("doneEvent")
	return nil
}

func (app App) publishEvent(c *gin.Context) error {
	log.Info("publishEvent")
	return nil
}

func (app App) PublishDoneEvent(c *gin.Context) error {
	log.Info("PublishDoneEvent")
	return nil
}

func (app App) playEvent(c *gin.Context) error {
	log.Info("playEvent")
	return nil
}

func (app App) playDoneEvent(c *gin.Context) error {
	log.Info("playDoneEvent")
	return nil
}

func (app App) updateEvent(c *gin.Context) error {
	log.Info("updateEvent")
	return nil
}
