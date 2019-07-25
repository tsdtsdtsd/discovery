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
	)

	flag.IntVar(&broadcastPort, "port", 1300, "broadcast port")
	flag.Parse()

	// Start
	discoveryServer := discovery.NewServer(broadcastPort)

	// Add some simple handlers and start listening.
	// The DefaultHandler could invoke automatic registration of clients.
	// Handlers will be executed concurrently in this POC, so don't expect ordered computing.
	discoveryServer.AddHandler(discoveryServer.DefaultHandler)
	discoveryServer.AddHandler(broadcastLogging)
	discoveryServer.Listen()

	// Block forever
	select {}
}

func broadcastLogging(incoming discovery.BroadcastMessage) {
	fmt.Printf("Incoming message from client with IP %s: ", incoming.RemoteAddr)
	fmt.Println(string(incoming.Data))
}
