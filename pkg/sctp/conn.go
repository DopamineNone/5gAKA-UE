package sctp

import (
	"github.com/free5gc/sctp"
	"syscall"
)

// ConnectionHandler must add `defer conn.Close()`
type ConnectionHandler interface {
	Handle(conn *Conn)
}

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

func (c *Conn) SetReadTimeout(milliseconds int) error {
	return c.conn.SetReadTimeout(syscall.Timeval{Sec: int64(milliseconds / 1000), Usec: int64((milliseconds % 1000) * 1000)})
}

func (c *Conn) SetWriteTimeout(milliseconds int) error {
	return c.conn.SetWriteTimeout(syscall.Timeval{Sec: int64(milliseconds / 1000), Usec: int64((milliseconds % 1000) * 1000)})
}

func (c *Conn) Stop() {
	if c.conn != nil {
		_ = c.conn.Close()
	}
}
