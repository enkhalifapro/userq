package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"

	"gopkg.in/mgo.v2/bson"

	"github.com/enkhalifapro/userq/msgq"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", viper.GetString("db.uri"))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

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
		// create in-memory datastore
		pool := newPool()
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
			for k, v := range msgMap {
				key := fmt.Sprintf("msg:%s:%s", messageID, k)
				c.Send("SET", key, v)
			}
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
