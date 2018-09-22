package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/enkhalifapro/userq/db"

	"gopkg.in/mgo.v2/bson"

	"github.com/enkhalifapro/userq/msgq"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func toMsgObj(msgStr string) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	err := json.Unmarshal([]byte(msgStr), &obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

var runMsgQ = &cobra.Command{
	Use:   "msgqsrv",
	Short: "run pubsub message queue listner",
	Run: func(cmd *cobra.Command, args []string) {
		// create redis db connection
		pool := db.NewPool()
		c := pool.Get()
		defer c.Close()

		// msgQ listener
		msgq.Listen(viper.GetString("msgqurl"), func(msg string, err error) {
			// save received messages into memoryQ
			messageID := bson.NewObjectId().Hex()
			msgMap, err := toMsgObj(msg)
			if err != nil {
				log.Fatalf(err.Error())
			}

			// save with key schema ex. `msg:$ID:$FIELD`
			// fullname
			key := fmt.Sprintf("msg:%s:%s", messageID, "fullname")
			c.Send("SET", key, fmt.Sprintf("%s %s", msgMap["FirstName"], msgMap["LastName"]))
			// address
			key = fmt.Sprintf("msg:%s:%s", messageID, "address")
			c.Send("SET", key, msgMap["Address"])
			// gender
			key = fmt.Sprintf("msg:%s:%s", messageID, "gender")
			c.Send("SET", key, msgMap["Gender"])
			// timestamp
			key = fmt.Sprintf("msg:%s:%v", messageID, "timestamp")
			c.Send("SET", key, msgMap["Timestamp"])

			// flush all
			c.Flush()
			c.Receive()          // reply from SET
			_, err = c.Receive() // reply from GET
			if err != nil {
				log.Fatalln(err.Error())
			}
		})
	},
}

func init() {
	RootCmd.AddCommand(runMsgQ)
}
