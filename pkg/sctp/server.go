package sctp

import (
	"github.com/free5gc/sctp"
)

type Server struct {
	Listener *sctp.SCTPListener
	PPID     uint32
	ConnPool []*Conn
}

func NewServer(port int, ppid uint32) (*Server, error) {
	listener, err := sctp.ListenSCTP("sctp", &sctp.SCTPAddr{Port: port})
	if err != nil {
		return &Server{}, err
	}
	return &Server{
		Listener: listener,
		PPID:     ppid,
	}, nil
}

func (s *Server) Listen(timeout int, handler ConnectionHandler) {
	if timeout == 0 {
		timeout = -1
	}
	for {
		if conn, err := s.Listener.Accept(timeout); err == nil {
			newConn := NewConnection(conn.(*sctp.SCTPConn), s.PPID)
			err = newConn.conn.SubscribeEvents(sctp.SCTP_EVENT_DATA_IO)
			if err != nil {
				continue
			}
			s.ConnPool = append(s.ConnPool, newConn)
			go handler.Handle(newConn)
		}
	}
}

func (s *Server) Stop() {
	if s.Listener != nil {
		_ = s.Listener.Close()
	}
	if s.ConnPool != nil {
		for _, conn := range s.ConnPool {
			conn.Stop()
		}
	}
}
