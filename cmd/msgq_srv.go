package cmd

import (
	"fmt"
	"log"

	"github.com/tidwall/buntdb"
	"gopkg.in/mgo.v2/bson"

	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/pull"
	"github.com/go-mangos/mangos/transport/all"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runMsgQ = &cobra.Command{
	Use:   "msgqsrv",
	Short: "run pubsub message queue listner",
	Run: func(cmd *cobra.Command, args []string) {
		// create in-memory datastore
		memQDB, err := buntdb.Open(":memory:")
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
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
			messageID := bson.NewObjectId().Hex()
			err := memQDB.Update(func(tx *buntdb.Tx) error {
				_, _, err := tx.Set(messageID, string(serverMsg.Body), nil)
				return err
			})
			if err != nil {
				fmt.Printf("\nMessage receive failed: %v", err)
				return
			}

		}
	},
}

func init() {
	RootCmd.AddCommand(runMsgQ)
}