package route

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		fmt.Printf("[alert] parse req failed, err=[%v]\n", err)
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}

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

	body, _ := json.Marshal(smsReq)
	reqHttp, err := http.NewRequest("POST", config.GetConfig().Alert.Addr, bytes.NewBuffer(body))
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
	}
	reqHttp.Header.Set("Content-Type", "application/json")
	reqHttp.Header.Set("X-Secret", config.GetConfig().Alert.Secret)

	client := &http.Client{}
	resp, err := client.Do(reqHttp)
	if err != nil {
		fmt.Printf("[alert] request sms service fail, err=[%v], req=[%v]\n", err, reqHttp)
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	defer resp.Body.Close()
	c.JSON(200, gin.H{
		"err": "ok",
	})

	rspBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("[alert] ok send req=[%s], rsp=[%s]\n", string(body), string(rspBody))
}
