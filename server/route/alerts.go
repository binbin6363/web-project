package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AlertsHandler(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("alert data: %v\n", string(data))
	c.JSON(200, nil)
}
