package api

import (
	"live/internal/exception"
	log "live/internal/helper"
	"live/internal/response"
	"live/internal/service"
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
		event.POST("on_publish_done", exception.ErrHandler(app.publishDoneEvent))
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
	log.Infof("connectEvent()")

	var event service.RtmpEvent
	if err := c.ShouldBind(&event); err != nil {
		log.Errorf("connectEvent(), bind param error: %v", err)
		return nil
	}

	key := c.DefaultPostForm("name", "")
	log.Infof("connectEvent(), key:%v, event params: %+v", key, event)

	c.JSON(http.StatusNoContent, nil)

	return nil
}

func (app App) doneEvent(c *gin.Context) error {
	log.Info("doneEvent()")
	var event service.RtmpEvent
	if err := c.ShouldBind(&event); err != nil {
		log.Errorf("doneEvent(), bind param error: %v", err)
		return nil
	}

	key := c.DefaultPostForm("name", "")
	log.Infof("doneEvent(), key:%v, event params: %+v", key, event)

	c.JSON(http.StatusNoContent, nil)

	return nil
}

func (app App) publishEvent(c *gin.Context) error {
	log.Info("publishEvent()")
	var event service.RtmpEvent
	if err := c.ShouldBind(&event); err != nil {
		log.Errorf("publishEvent(), bind param error: %v", err)
		return nil
	}

	key := c.DefaultPostForm("name", "")
	log.Infof("publishEvent(), key:%v, event params: %+v", key, event)

	c.JSON(http.StatusNoContent, nil)

	return nil
}

func (app App) publishDoneEvent(c *gin.Context) error {
	log.Info("publishDoneEvent()")
	var event service.RtmpEvent
	if err := c.ShouldBind(&event); err != nil {
		log.Errorf("publishDoneEvent(), bind param error: %v", err)
		return nil
	}

	key := c.DefaultPostForm("name", "")
	log.Infof("publishDoneEvent(), key:%v, event params: %+v", key, event)

	c.JSON(http.StatusNoContent, nil)

	return nil
}

func (app App) playEvent(c *gin.Context) error {
	log.Info("playEvent()")
	var event service.RtmpEvent
	if err := c.ShouldBind(&event); err != nil {
		log.Errorf("playEvent(), bind param error: %v", err)
		return nil
	}

	key := c.DefaultPostForm("name", "")
	log.Infof("playEvent(), key:%v, event params: %+v", key, event)

	c.JSON(http.StatusNoContent, nil)

	return nil
}

func (app App) playDoneEvent(c *gin.Context) error {
	log.Info("playDoneEvent")
	var event service.RtmpEvent
	if err := c.ShouldBind(&event); err != nil {
		log.Errorf("playDoneEvent(), bind param error: %v", err)
		return nil
	}

	key := c.DefaultPostForm("name", "")
	log.Infof("playDoneEvent(), key:%v, event params: %+v", key, event)

	c.JSON(http.StatusNoContent, nil)

	return nil
}

func (app App) updateEvent(c *gin.Context) error {
	log.Info("updateEvent()")
	var event service.RtmpEvent
	if err := c.ShouldBind(&event); err != nil {
		log.Errorf("updateEvent(), bind param error: %v", err)
		return nil
	}

	key := c.DefaultPostForm("name", "")
	log.Infof("updateEvent(), key:%v, event params: %+v", key, event)

	c.JSON(http.StatusNoContent, nil)

	return nil
}
