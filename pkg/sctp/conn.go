package sctp

import (
	"github.com/free5gc/sctp"
)

type Conn struct {
	conn *sctp.SCTPConn
	ppid uint32
}

func NewConnection(conn *sctp.SCTPConn, ppid uint32) *Conn {
	return &Conn{
		conn: conn,
		ppid: ppid,
	}
}

func (c *Conn) Send(data []byte, stream uint16) (int, error) {
	info := &sctp.SndRcvInfo{
		Stream: stream,
		PPID:   c.ppid,
	}
	return c.conn.SCTPWrite(data, info)
}

func (c *Conn) Receive(buf []byte) (n int, info *sctp.SndRcvInfo, notification sctp.Notification, err error) {
	n, info, notification, err = c.conn.SCTPRead(buf)
	return
}

func (c *Conn) Destruct() {
	if c.conn != nil {
		_ = c.conn.Close()
	}
}
