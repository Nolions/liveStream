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
	streamKey := helper.GenerateGUID()
	expiration := time.Now().Add(time.Duration(live.Conf.Expire) * time.Hour)

	return Stream{
		Push: Push{
			RTMP: live.rtmpPushUrl(streamKey, uid, expiration),
		},
		Pull: Pull{
			HLS: live.hlsPullUrl(streamKey, uid, expiration),
		},
		ExpirationTime: time.Now().Add(24 * time.Hour).UTC(),
	}, nil
}

// 產生以下格式rtmp推流url
// rtmp://localhost:1935/live/123456789?u=uid=md5(key+uuid+uid+hex(time))&t=hex(time)
func (live *Live) rtmpPushUrl(streamKey, uid string, expiration time.Time) string {
	timestamp := fmt.Sprintf("%x", expiration.Unix())

	elems := []string{
		live.Conf.Secret,
		streamKey,
		uid,
		timestamp,
	}
	secret := helper.EncryptMD5(elems)

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

// 產生以下格式HLS拉流url
// http://localhost:8080/hls/123456789.m3u8?u=uid=md5(key+uuid+uid+hex(time))&t=hex(time)
func (live *Live) hlsPullUrl(streamKey, uid string, expiration time.Time) string {
	timestamp := fmt.Sprintf("%x", expiration.Unix())
	elems := []string{
		live.Conf.Secret,
		streamKey,
		uid,
		timestamp,
	}
	secret := helper.EncryptMD5(elems)

	url := fmt.Sprintf("%s/%s?u=%s&s=%s&t=%s",
		live.Conf.HLSUrl,
		streamKey,
		uid,
		secret,
		timestamp,
	)

	return url
}
