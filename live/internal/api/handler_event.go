package api

import (
	log "live/internal/helper"
	"live/internal/response"
	"live/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 建立連線事件
func (app App) connectEvent(c *gin.Context) error {
	log.Infof("connectEvent()")

	var event service.RtmpEvent
	if err := c.ShouldBind(&event); err != nil {
		log.Errorf("connectEvent(), bind param error: %v", err)
		return nil
	}

	key := c.DefaultPostForm("name", "")
	log.Infof("connectEvent(), key:%v, event params: %+v", key, event)

	if app.Live.ConnectEvent(&event) {
		c.JSON(http.StatusNoContent, nil)
	} else {
		response.HandleBadRequest(c, "No allow connect")
	}

	return nil
}

// 開始推流事件
func (app App) publishEvent(c *gin.Context) error {
	log.Info("publishEvent()")
	var event service.RtmpEvent
	if err := c.ShouldBind(&event); err != nil {
		log.Errorf("publishEvent(), bind param error: %v", err)
		return nil
	}

	key := c.DefaultPostForm("name", "")
	log.Infof("publishEvent(), key:%v, event params: %+v", key, event)
	// TODO 根據 RTMP url進行鑑權
	if app.Live.PublishEvent() {
		c.JSON(http.StatusNoContent, nil)
	} else {
		response.HandleBadRequest(c, "no allow publish")
	}

	return nil
}

// 推流狀態更新
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

// 結束連線
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

// 結束推流
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
