package manager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"linkwan.cn/linkExporter/src/handler/linkStateUpdate"

	log "linkwan.cn/linkExporter/src/log"
	iotController "linkwan.cn/linkExporter/src/proto/natsMessage"

	"github.com/gogo/protobuf/proto"
	nats "github.com/nats-io/nats.go"
)

type consumer struct {
	nc *nats.Conn
}

// Start method starts to consume normalized messages received from NATS.
func Start(nc *nats.Conn, queue string) error {
	c := consumer{
		nc: nc,
	}

	_, err := nc.QueueSubscribe("channelState", queue, c.consume)
	return err
}

func (c *consumer) consume(m *nats.Msg) {
	log.Infof("recv nats msg: %s", string(m.Data))
	msg := &iotController.RawMessage{}
	if err := proto.Unmarshal(m.Data, msg); err != nil {
		log.Error(fmt.Sprintf("Failed to unmarshal received message: %s", err))
		return
	}

	var iotMsg Message
	decoder := json.NewDecoder(bytes.NewReader(msg.GetPayload()))
	decoder.UseNumber()

	if err := decoder.Decode(&iotMsg); err != nil {
		log.Error(fmt.Sprintf("Failed to unmarshal message unit: %s", err))
		return
	}
	if !validateMessage(&iotMsg) {
		log.Error("Failed to validate tunnel message: %v", iotMsg)
		return
	}

	switch iotMsg.Topic {

	case ReportTunnelState:
		log.Info("state %v", iotMsg)
		linkStateUpdate.Handle(iotMsg.MsgId, msg.GetChannel(), iotMsg.Payload)
	}
	log.Info(iotMsg)
}

func validateMessage(tunnelMsg *Message) bool {
	if tunnelMsg.MsgId == "" {
		return false
	}
	if tunnelMsg.Topic == "" {
		return false
	}
	return true
}
