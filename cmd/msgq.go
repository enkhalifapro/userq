package cmd

import (
	"fmt"
	"log"

	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/pull"
	"github.com/go-mangos/mangos/transport/all"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runMsgQ = &cobra.Command{
	Use:   "msgq",
	Short: "run pubsub message queue",
	Run: func(cmd *cobra.Command, args []string) {
		// msgQ listener
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

		if err = pullServer.Listen(viper.GetString("msgqurl")); err != nil {
			log.Fatalf("\nMSGQ Server listen failed: %v", err)
			return
		}
		log.Printf("Listening and serving MsgQ on %s", viper.GetString("msgqurl"))

		close(pullServerReady)

		for {
			// fmt.Println(fmt.Sprintf("\nmsgQ listening at %v", "tcp://127.0.0.1:7000"))
			if serverMsg, err = pullServer.RecvMsg(); err != nil {
				log.Fatalf("\nServer receive failed: %v", err)
			}
			fmt.Println("in rec vvvvv")
			fmt.Println(serverMsg.Body)
		}
	},
}

func init() {
	RootCmd.AddCommand(runMsgQ)
}
