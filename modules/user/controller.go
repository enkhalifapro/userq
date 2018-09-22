package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service *Service `inject:""`
}

func (c *Controller) Test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Ok...")
}
