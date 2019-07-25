package main

import (
	"flag"
	"fmt"

	"github.com/tsdtsdtsd/discovery"
)

func main() {

	// CLI args
	var (
		broadcastPort int
		broadcastIP   string
		broadcastMsg  string
	)

	flag.IntVar(&broadcastPort, "port", 1300, "broadcast port")
	flag.StringVar(&broadcastIP, "ip", "", "broadcast IP")
	flag.StringVar(&broadcastMsg, "msg", "It's me, Mario!", "broadcast message")
	flag.Parse()

	// Contruct client and send broadcast message
	discoveryClient := discovery.NewClient(broadcastPort)
	err := discoveryClient.Broadcast(broadcastIP, broadcastMsg)

	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
}
