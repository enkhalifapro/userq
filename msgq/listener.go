package msgq

import (
	"log"

	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/pull"
	"github.com/go-mangos/mangos/transport/all"
)

// RecvMsgFunc a callback function for nanoMsg receive message event
type RecvMsgFunc func(msg string, err error)

// Listen nanoMsg protocol socket
func Listen(url string, fn RecvMsgFunc) {
	// 1. Create Pull Server
	pullServerReady := make(chan struct{})
	pullServer, err := pull.NewSocket()
	if err != nil {
		log.Fatalf("Cannot create socket: %s\n", err.Error())
	}
	defer pullServer.Close()

	all.AddTransports(pullServer)

	// 2. Run Pull Server
	var serverMsg *mangos.Message

	if err = pullServer.Listen(url); err != nil {
		log.Fatalf("\nMSGQ Server listen failed: %v", err)
		return
	}
	log.Printf("Listening and serving MsgQ on %s", url)

	close(pullServerReady)

	for {
		serverMsg, err = pullServer.RecvMsg()
		fn(string(serverMsg.Body), err)
	}
}
