package route

import (
	"github.com/gin-gonic/gin"
)

func ConfigHandler(c *gin.Context) {

	c.JSON(200, Result{
		Data:       "config",
		ResultCode: 200,
		Message:    "ok",
	})
}

func ConfigDelHandler(c *gin.Context) {

	c.JSON(200, Result{
		Data:       "config del",
		ResultCode: 200,
		Message:    "ok",
	})
}
