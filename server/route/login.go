package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loginHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"path": c.FullPath(),
	})
}