package nts

const (
	TimeExpired int = iota

	ConnectionSetup
	ConnectionShutdown
	ReceiveMessage
	SendMessage

	UplinkDelivery
	DownlinkDelivery
)

// Message base message struct
type Message struct {
	MessageType int
	PDU         []byte
}
