package route

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/binbin6363/web-project/server/config"
	"github.com/binbin6363/web-project/server/model"
	"github.com/gin-gonic/gin"
)

func State2Action(state string) string {
	switch state {
	case "alerting":
		return "firing"
	default:
		return "resolved"
	}
}

// AlertsHandler .
func AlertsHandler(c *gin.Context) {
	req := new(model.AlertRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		fmt.Printf("parse req failed, err:%v\n", err)
		c.JSON(200, nil)
		return
	}

	fmt.Printf("recv: %v\n", req)
	meta := fmt.Sprintf("match count:%d", len(req.EvalMatches))

	smsReq := &model.SmsAlertReq{
		Ip:           req.RuleName,
		SourceTime:   time.Now().Format("2006-01-02 15:04:05"),
		AlarmType:    req.RuleName,
		Level:        "warning",
		AlarmName:    req.Title,
		AlarmContent: req.Message,
		MetaInfo:     meta,
		Action:       State2Action(req.State),
	}

	fmt.Printf("smsReq: %v, alert cfg:%v\n", smsReq, config.GetConfig().Alert)
	body, _ := json.Marshal(smsReq)
	reqHttp, err := http.NewRequest("POST", config.GetConfig().Alert.Addr, bytes.NewBuffer(body))
	if err != nil {
		fmt.Printf("make sms req fail, err:%v, body:%s\n", err, string(body))
		c.JSON(500, nil)
	}
	reqHttp.Header.Set("Content-Type", "application/json")
	reqHttp.Header.Set("X-Secret", config.GetConfig().Alert.Secret)

	fmt.Printf("after Set Header:%v, body:%s\n", reqHttp.Header, string(body))
	client := &http.Client{}
	resp, err := client.Do(reqHttp)
	if err != nil {
		fmt.Printf("req sms fail, err:%v, req:%v\n", err, reqHttp)
		c.JSON(500, resp)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("after client.Do:%v, body:%s\n", reqHttp.Header, string(body))
	c.JSON(200, nil)
}
