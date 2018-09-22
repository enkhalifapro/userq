package cmd

import (
	"log"

	"github.com/enkhalifapro/userq/server"
	"github.com/facebookgo/inject"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run server",
	Run: func(cmd *cobra.Command, args []string) {
		// Creates a gin router with default middleware:
		engine := gin.Default()

		// CORS config
		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		config.AllowCredentials = true
		config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Access-Control-Allow-Origin", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "accept", "origin", "Cache-Control", "X-Requested-With"}
		config.AllowMethods = []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"}

		engine.Use(cors.New(config))

		server := &server.Server{}

		graph := &inject.Graph{}
		err := graph.Provide(
			// Provide engine
			&inject.Object{Value: engine},
			&inject.Object{Value: server},
		)
		if err != nil {
			log.Fatal(err)
			return
		}

		if err := graph.Populate(); err != nil {
			log.Fatal(err)
			return
		}

		if err := server.Run(); err != nil {
			log.Fatal(err)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(runCmd)
}
