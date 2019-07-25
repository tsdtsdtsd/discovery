package discovery

import (
	"net"
	"time"
)

// BroadcastMessage contains details of a broadcast
type BroadcastMessage struct {
	Time       time.Time
	RemoteAddr net.IP
	Data       []byte
	Length     int
}
