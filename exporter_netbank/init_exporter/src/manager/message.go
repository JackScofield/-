package manager

const ReportTunnelState = "ReportTunnelState"

type Message struct {
	MsgId   string
	Topic   string
	ErrMsg  string
	Payload interface{}
}
