package sctp

import (
	"github.com/free5gc/sctp"
	"net"
)

type Client struct {
	remoteAddr *sctp.SCTPAddr
	localAddr  *sctp.SCTPAddr
	*Conn
}

func NewClient(ips []string, port, lPort int, ppid uint32) (*Client, error) {
	// initialize sctp addr
	addr := make([]net.IPAddr, len(ips))
	for i, ip := range ips {
		if ipAddr, err := net.ResolveIPAddr("ip", ip); err == nil {
			addr[i] = *ipAddr
		} else {
			return &Client{}, err
		}
	}

	conn := &Conn{ppid: ppid, conn: nil}
	return &Client{
		remoteAddr: &sctp.SCTPAddr{
			IPAddrs: addr,
			Port:    port,
		},
		localAddr: &sctp.SCTPAddr{
			Port: lPort,
		},
		Conn: conn,
	}, nil
}

func (c *Client) Connect() error {
	if conn, err := sctp.DialSCTP("sctp", c.localAddr, c.remoteAddr); err == nil {
		c.conn = conn
	} else {
		return err
	}

	return c.conn.SubscribeEvents(sctp.SCTP_EVENT_DATA_IO)
}
