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

type Config struct {
	IPs   []string
	Port  int
	LPort int
	PPID  uint32
}

func NewClient(cfg Config) (*Client, error) {
	// initialize sctp addr
	addr := make([]net.IPAddr, len(cfg.IPs))
	for i, ip := range cfg.IPs {
		if ipAddr, err := net.ResolveIPAddr("ip", ip); err == nil {
			addr[i] = *ipAddr
		} else {
			return &Client{}, err
		}
	}

	conn := &Conn{ppid: cfg.PPID, conn: nil}
	return &Client{
		remoteAddr: &sctp.SCTPAddr{
			IPAddrs: addr,
			Port:    cfg.Port,
		},
		localAddr: &sctp.SCTPAddr{
			Port: cfg.LPort,
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
