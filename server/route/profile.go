package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func profileHandler(c *gin.Context) {
	fmt.Printf("show profile req")
	c.JSON(200, Result{
		Data:          "profile",
		ResultCode:    200,
		Message:       "ok",
		NickName:      "polite",
		LoginUserName: "admin",
	})

}
