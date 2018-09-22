package server

import (
	"github.com/enkhalifapro/userq/modules/user"
)

// Controller collection struct for server.
type Controller struct {
	User *user.Controller `inject:""`
}
