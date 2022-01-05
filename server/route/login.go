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
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "content-type, token, x-requested-with")
	c.Header("Access-Control-Allow-Methods", "GET,HEAD,POST,PUT,DELETE,OPTIONS")
	c.Header("Access-Control-Max-Age", "3600")
	c.Header("Access-Control-Allow-Credentials", "true")

	c.JSON(200, gin.H{
		"Blog":   "www.flysnow.org",
		"wechat": "flysnow_org",
	})
}
