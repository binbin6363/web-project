package route

import (
	"github.com/gin-gonic/gin"
)


func Route(r *gin.Engine)  {
	// group: v1
	v1 := r.Group("/manage-api/v1")
	{
		v1.POST("/adminUser/login", loginHandler)
		v1.POST("/indexConfigs/delete", ConfigDelHandler)
		v1.POST("/carousels", CarouselsHandler)
		v1.POST("/indexConfigs", ConfigHandler)
		v1.POST("/logout", LogoutHandler)
	}
}
