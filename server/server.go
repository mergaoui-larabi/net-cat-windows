package server

import (
	"fmt"
	"net"
)

type Message struct {
	from    string
	content []byte
}

type Server struct {
	listenerAddr string
	listener     net.Listener
	quitch       chan struct{}
	msgch        chan Message
}

func NewServer(addr string) *Server {
	return &Server{
		listenerAddr: addr,
		quitch:       make(chan struct{}),
		msgch:        make(chan Message),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.listenerAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.listener = ln

	go s.AcceptLoop()
	<-s.quitch

	return nil
}

func (s *Server) AcceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		go s.ReadLoop(conn)
	}
}

func (s *Server) ReadLoop(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		s.msgch <- Message{
			from:    conn.RemoteAddr().String(),
			content: buffer[:n],
		}
		conn.Write([]byte("tank u for ur message\ngit a"))
	}
}

func (s *Server) Display() {
	for msg := range s.msgch {
		fmt.Printf("from:%s,msg:%s", msg.from, string(msg.content))
	}
}
