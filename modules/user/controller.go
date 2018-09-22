package user

import (
	"net/http"

	"github.com/enkhalifapro/userq/msgq"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service    *Service         `inject:""`
	MsgqHelper *msgq.MsgQHelper `inject:""`
}

// Post pushes user to msgQ
func (c *Controller) Post(ctx *gin.Context) {
	var user Model
	err := ctx.Bind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// send to msgQ
	err = c.MsgqHelper.Push(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
