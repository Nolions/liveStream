package service

import log "live/internal/helper"

type RtmpEvent struct {
	App      string `form:"app"`
	Addr     string `form:"addr"`
	ClientId string `form:"clientid"`
	Call     string `form:"call"`
	Name     string `form:"name"`
}

// ConnectEvent
// rtmp連線事件
// step1. IP 黑白名單驗證
// step2. 檢查APP名稱的路由
func (live *Live) ConnectEvent(event *RtmpEvent) bool {
	log.Infof("ConnectEvent(), New connection from %s for app: %s", event.Addr, event.App)
	// IP 黑白名單驗證
	if isIpInBlacklist(event.Addr) {
		return false
	}

	// 檢查APP名稱的路由
	if !verifyAppName(event.App) {
		return false
	}

	return true
}

func (live *Live) PublishEvent() bool {
	return true
}

func verifyAppName(app string) bool {
	appNames := []string{"live"} // TODO 後續可以改用從DB拿掉的APP
	for _, name := range appNames {
		if name == app {
			return true
		}
	}

	return false
}

// 檢查IP是否在黑名單內
func isIpInBlacklist(ip string) bool {
	bIps := []string{} // TODO 後續可以改用從DB拿掉的拿出黑名單的IP
	for _, bip := range bIps {
		if ip == bip {
			return true
		}
	}
	return false
}
