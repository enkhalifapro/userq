package msgq

import (
	"encoding/json"

	"nanomsg.org/go-mangos/protocol/push"
	"nanomsg.org/go-mangos/transport/all"
)

// MsgQHelper contains functions to interact with msgQ
type MsgQHelper struct {
	msgQURL string
}

// NewMsgQHelper creates new msgQHelper instance
func NewMsgQHelper(msgQURL string) *MsgQHelper {
	return &MsgQHelper{msgQURL: msgQURL}
}

// Push a message
func (s *MsgQHelper) Push(msg interface{}) error {
	json, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	pushSocket, err := push.NewSocket()
	defer pushSocket.Close()
	all.AddTransports(pushSocket)

	if err = pushSocket.Dial(s.msgQURL); err != nil {
		return err
	}
	err = pushSocket.Send(json)
	return err
}
