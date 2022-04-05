package route

import (
	"github.com/gin-gonic/gin"
)

type Result struct {
	Data          string `json:"data,omitempty"`
	ResultCode    int    `json:"resultCode,omitempty"`
	Message       string `json:"message,omitempty"`
	NickName      string `json:"nickName,omitempty"`
	LoginUserName string `json:"loginUserName,omitempty"`
}

func Route(r *gin.Engine) {
	// group: v1
	v1 := r.Group("/manage-api/v1")
	{
		v1.POST("/adminUser/login", loginHandler)
		v1.GET("/adminUser/profile", profileHandler)
		v1.POST("/indexConfigs/delete", ConfigDelHandler)
		v1.GET("/carousels", CarouselsHandler)
		v1.GET("/indexConfigs", ConfigHandler)
		v1.DELETE("/logout", LogoutHandler)
		v1.GET("/orders", GetOrderHandler)
	}
}
