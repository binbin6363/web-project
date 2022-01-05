package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func loginHandler(c *gin.Context) {
	fmt.Printf("show login req")

	c.JSON(200, Result{
		Data:       uuid.Must(uuid.NewV4()).String(),
		ResultCode: 200,
		Message:    "ok",
	})

}

func LogoutHandler(c *gin.Context) {
	fmt.Printf("show logout req")

	c.JSON(200, Result{
		Data:       "logout",
		ResultCode: 200,
		Message:    "ok",
	})
}
