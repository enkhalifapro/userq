package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service *Service `inject:""`
}

func (c *Controller) Test(ctx *gin.Context) {
	time.Now().Unix()
	ctx.JSON(http.StatusOK, "Ok...")
}
