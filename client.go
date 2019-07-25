package discovery

import (
	"fmt"
	"net"
)

type Client struct {
	Port int
}

func NewClient(port int) (c *Client) {
	c = &Client{
		Port: port,
	}

	return c
}

func (c *Client) Broadcast(IP, message string) error {

	IPv4 := net.ParseIP(IP)
	if IPv4 == nil {
		return fmt.Errorf("invalid IP address for broadcast")
	}

	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   IPv4,
		Port: c.Port,
	})

	if err != nil {
		// Couldn't dial UDP connection
		return err
	}

	_, err = socket.Write([]byte(message))
	if err != nil {
		// Socket write failed.
		// Non-failure does NOT mean, that the server received our message!
		return err
	}

	return socket.Close()
}
