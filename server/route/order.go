package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type OrderItem struct {
	OrderNo     string `json:"orderNo,omitempty"`
	OrderId     string `json:"orderId,omitempty"`
	TotalPrice  string `json:"totalPrice,omitempty"`
	OrderStatus string `json:"orderStatus,omitempty"`
	PayType     string `json:"payType,omitempty"`
	CreateTime  string `json:"createTime,omitempty"`
}

type OrderInfo struct {
	ResultCode int         `json:"resultCode,omitempty"`
	Message    string      `json:"message,omitempty"`
	TotalCount int         `json:"totalCount,omitempty"`
	OrderList  []OrderItem `json:"list,omitempty"`
	CurrPage   int         `json:"currPage,omitempty"`
}

func GetOrderHandler(c *gin.Context) {
	fmt.Printf("show order req")

	order := &OrderInfo{
		TotalCount: 1,
		CurrPage:   1,
		ResultCode: 200,
		Message:    "ok",
	}
	orderItem1 := OrderItem{
		OrderNo:     uuid.Must(uuid.NewV4()).String(),
		OrderId:     "1",
		TotalPrice:  "100",
		OrderStatus: "1",
		PayType:     "1",
		CreateTime:  "20220301",
	}
	orderItem2 := OrderItem{
		OrderNo:     uuid.Must(uuid.NewV4()).String(),
		OrderId:     "2",
		TotalPrice:  "100",
		OrderStatus: "0",
		PayType:     "1",
		CreateTime:  "20220311",
	}
	orderItem3 := OrderItem{
		OrderNo:     uuid.Must(uuid.NewV4()).String(),
		OrderId:     "3",
		TotalPrice:  "100",
		OrderStatus: "1",
		PayType:     "0",
		CreateTime:  "20220321",
	}
	orderItem4 := OrderItem{
		OrderNo:     uuid.Must(uuid.NewV4()).String(),
		OrderId:     "4",
		TotalPrice:  "100",
		OrderStatus: "1",
		PayType:     "1",
		CreateTime:  "20220401",
	}
	order.OrderList = append(order.OrderList, orderItem1)
	order.OrderList = append(order.OrderList, orderItem2)
	order.OrderList = append(order.OrderList, orderItem3)
	order.OrderList = append(order.OrderList, orderItem4)
	c.JSON(200, order)

}
