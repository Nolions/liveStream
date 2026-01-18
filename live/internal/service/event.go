package service

type RtmpEvent struct {
	App      string `form:"app"`
	Addr     string `form:"addr"`
	ClientId string `form:"clientid"`
	Call     string `form:"call"`
	Name     string `form:"name"`
}
