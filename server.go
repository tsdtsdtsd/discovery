package discovery

import (
	"fmt"
	"net"
	"time"
)

// Server defines the listening part for discovery
type Server struct {
	Port     int
	handlers []func(incoming BroadcastMessage)
}

// NewServer returns a default discovery.Server object.
func NewServer(port int) (s *Server) {
	s = &Server{
		Port: port,
	}

	return s
}

// Listen stars listening for UDP broadcasts on the network
func (s *Server) Listen() error {

	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: s.Port,
	})

	if err != nil {
		// Couldn't listen on UDP connection
		return err
	}

	go func() {
		for {
			data := make([]byte, 4096)
			read, remoteAddr, err := socket.ReadFromUDP(data)

			if err != nil {
				fmt.Printf("error while reading from UDP connection: %s\n", err.Error())
				continue
			}

			message := BroadcastMessage{
				Time:       time.Now(),
				RemoteAddr: remoteAddr.IP,
				Data:       data,
				Length:     read,
			}

			s.invokeHandlers(message)
		}
	}()

	fmt.Println("Started listening ...")
	return nil
}

// AddHandler adds a handler to the stack.
// Handlers on the stack will be called on every incoming broadcast message.
func (s *Server) AddHandler(h func(incoming BroadcastMessage)) {
	s.handlers = append(s.handlers, h)
}

// DefaultHandler could invoke the automatic registration of clients
func (s *Server) DefaultHandler(incoming BroadcastMessage) {

}

func (s *Server) invokeHandlers(msg BroadcastMessage) {

	if len(s.handlers) <= 0 {
		return
	}

	for _, handler := range s.handlers {
		go handler(msg)
	}
}
