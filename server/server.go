// --
// --Add deployment
// Prepare deployment model or use "GET /api/v1/sample" to obtain one
// Add all required inputs and any other optional data
// "POST /api/v1/config" add deployment data to DB
// "POST /api/v1/automation/:shortname" to run add deployment against kubernetes and create desired "pod" (edited)
//
// --
// GET /api/v1/sample to get sample data of deployment
// POST /api/v1/config add config to DB
// POST /api/v1/automation add config to DB
// GET /api/v1/config/:name get deployment config by short name
// GET /api/v1/configs lists all saved configurations

package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Server which handles API requests
type Server struct {
	Engine     *gin.Engine `inject:""`
	Controller *Controller `inject:""`
}

// Run server
func (s *Server) Run() error {
	// deployment APIs
	s.Engine.GET("/", s.Controller.User.Test)

	return s.Engine.Run(fmt.Sprintf("%v:%v", viper.GetString("host"), viper.GetString("port")))
}
