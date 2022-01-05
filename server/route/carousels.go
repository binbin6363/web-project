package route

import (
	"github.com/gin-gonic/gin"
)

func CarouselsHandler(c *gin.Context) {

	c.JSON(200, Result{
		Data:       "logout",
		ResultCode: 200,
		Message:    "ok",
	})
}
