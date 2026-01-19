package service

import (
	"context"
	"fmt"
	"live/conf"
	"live/internal/helper"
	"time"
)

type Stream struct {
	Push           Push      `json:"push"`
	Pull           Pull      `json:"pull"`
	ExpirationTime time.Time `json:"expiration_time"`
}

type Push struct {
	RTMP string `json:"rtmp"`
}

type Pull struct {
	HLS string `json:"hls"`
}

type Live struct {
	Ctx  context.Context
	Conf conf.Live
}

func NewLive(ctx context.Context, conf conf.Live) *Live {
	return &Live{
		Ctx:  ctx,
		Conf: conf,
	}
}

func (live *Live) GetStreamDate() (Stream, error) {
	uid := "1234567890"

	return Stream{
		Push: Push{
			RTMP: live.genRtmpPushUrl(uid),
		},
		Pull: Pull{
			HLS: "456",
		},
		ExpirationTime: time.Now().Add(24 * time.Hour).UTC(),
	}, nil
}

// 產生以下格式rtmp推流url
// rtmp://localhost:1935/live/<uuid>?secret=md5(key+uuid+uid+hex(time))&time=hex(time)
func (live *Live) genRtmpPushUrl(uid string) string {
	streamKey := helper.GenerateGUID()

	unixTime := time.Now().Add(24 * time.Hour).Unix()
	timestamp := fmt.Sprintf("%x", unixTime)

	elems := []string{
		live.Conf.Secret,
		streamKey,
		uid,
		timestamp,
	}
	secret := helper.SecretMD5(elems)

	url := fmt.Sprintf("%s/%s/%s?u=%s&s=%s&t=%s",
		live.Conf.RtmpUrl,
		live.Conf.AppName,
		streamKey,
		uid,
		secret,
		timestamp,
	)

	return url
}
