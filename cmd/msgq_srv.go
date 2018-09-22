package cmd

import (
	"fmt"
	"log"

	"github.com/carlescere/scheduler"
	"github.com/enkhalifapro/userq/msgq"
	"github.com/tidwall/buntdb"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/* func x(memQDB *buntdb.DB) {
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
		fmt.Println(string(serverMsg.Body))
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
} */

var runMsgQ = &cobra.Command{
	Use:   "msgqsrv",
	Short: "run pubsub message queue listner",
	Run: func(cmd *cobra.Command, args []string) {
		// create in-memory datastore
		memQDB, err := buntdb.Open(":memory:")
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}

		// save pending messages into redis
		pendingMemQ := func() {
			memQDB.View(func(tx *buntdb.Tx) error {
				count, _ := tx.Len()
				fmt.Printf("%v Pending messages in msgQ \n", count)
				return nil
			})

		}
		scheduler.Every(1).Seconds().Run(pendingMemQ)

		// msgQ listener
		fmt.Println(viper.GetString("msgqurl"))
		msgq.Listen(viper.GetString("msgqurl"), func(msg string, err error) {
			fmt.Println("in bbbbbb")
			fmt.Println(err)
			fmt.Println(msg)
		})
	},
}

func init() {
	RootCmd.AddCommand(runMsgQ)
}
