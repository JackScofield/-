package linkStateUpdate

import (
	"encoding/json"
	"linkwan.cn/linkExporter/src/cache"
	"time"

	"linkwan.cn/linkExporter/src/exporter"
	"linkwan.cn/linkExporter/src/log"
)

type reportTunnelState struct {
	LinkName           string
	MasterPop          string
	SlavePop           string
	Wan                []cache.Wan_t
	ConfigVersion      int64
	TunnelRouteVersion int64
}

func Handle(msgId, channelId string, reqPayloadObj interface{}) {
	bytes, err := json.Marshal(reqPayloadObj)
	log.Infof("channelId: %s", channelId)
	if err != nil {
		log.Error(err)
		return
	}
	var reqPayload reportTunnelState
	err = json.Unmarshal(bytes, &reqPayload)
	if err != nil {
		log.Error(err)
		return
	}
	log.Infof("linkName: %s", reqPayload.LinkName)

	channelStateCache := cache.GetChannelStateCache()
	slaveState, masterState := cache.Unknown, cache.Unknown

	for _, wan := range reqPayload.Wan {
		if wan.WanType == cache.Slave {
			slaveState = wan.WanState
		} else if wan.WanType == cache.Master {
			masterState = wan.WanState
		}
	}
	if _, exist := channelStateCache.Load(reqPayload.LinkName); !exist {
		if err := exporter.Register(reqPayload.LinkName); err != nil {
			log.Error("prometheus register error: ", err)
		}
	}

	channelStateCache.Store(reqPayload.LinkName, &cache.ChannelState_t{
		LastReportTime:    time.Now(),
		SlaveTunnelState:  slaveState,
		MasterTunnelState: masterState,
	})

}
