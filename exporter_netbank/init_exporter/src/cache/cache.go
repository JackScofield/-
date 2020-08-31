package cache

import (
	"sync"
	"time"
)

var (
	channelStataCache *sync.Map
	cacheOnce         sync.Once
)

type ChannelState_t struct {
	LastReportTime    time.Time
	SlaveTunnelState  WanState_t
	MasterTunnelState WanState_t
}

func GetChannelStateCache() *sync.Map {
	if channelStataCache == nil {
		cacheOnce.Do(func() {
			channelStataCache = new(sync.Map)
		})
	}
	return channelStataCache
}

type WanType_t string

const (
	Master WanType_t = "0"
	Slave  WanType_t = "1"
)

type WanState_t string

const (
	Up      WanState_t = "0"
	Down    WanState_t = "1"
	Unknown WanState_t = "2"
)

type Wan_t struct {
	TunnelId string
	WanType  WanType_t
	WanState WanState_t
}
