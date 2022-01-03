package route

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func loginHandler(c *gin.Context)  {
	fmt.Printf("show login req")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "content-type, token, x-requested-with")
	c.Header("Access-Control-Allow-Methods", "GET,HEAD,POST,PUT,DELETE,OPTIONS")
	c.Header("Access-Control-Max-Age", "3600")
	c.Header("Access-Control-Allow-Credentials", "true")


	c.JSON(200, gin.H{
		"Blog":"www.flysnow.org",
		"wechat":"flysnow_org",
	})
}

func LogoutHandler(c *gin.Context)  {
	fmt.Printf("show logout req")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "content-type, token, x-requested-with")
	c.Header("Access-Control-Allow-Methods", "GET,HEAD,POST,PUT,DELETE,OPTIONS")
	c.Header("Access-Control-Max-Age", "3600")
	c.Header("Access-Control-Allow-Credentials", "true")


	c.JSON(200, gin.H{
		"Blog":"www.flysnow.org",
		"wechat":"flysnow_org",
	})
}