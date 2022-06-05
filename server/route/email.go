package route

import (
	"fmt"
	"log"
	"time"

	"github.com/binbin6363/web-project/server/middleware"
	"github.com/binbin6363/web-project/server/model"
	"github.com/gin-gonic/gin"
)

// 创建email
func addEmail(c *gin.Context) {
	req := &model.AddEmailReq{}
	if err := c.ShouldBind(req); err != nil {
		log.Printf("bind req failed, err:%v", err)
		c.JSON(200, Result{
			Data:       nil,
			ResultCode: 500,
			Message:    "param error",
		})
		return
	}

	fmt.Printf("show add email req:%v", req)
	rsp := &model.AddEmailRsp{}
	the_time, _ := time.ParseInLocation("2006-01-02T15:04:05.000Z", req.TriggleDate, time.Local)

	emailInfo := &model.EmailTask{
		Subject:     req.Subject,
		Body:        req.Body,
		Creator:     req.Creator,
		Receiver:    req.Receiver,
		Public:      req.Public,
		Mode:        req.Mode,
		State:       0,
		TriggleTime: int(the_time.Unix()),
		CreateTime:  int(time.Now().Unix()),
		UpdateTime:  int(time.Now().Unix()),
	}
	fmt.Printf("show add email emailInfo:%v", emailInfo)
	if id, err := middleware.GetDao().CreateEmailTask(c.Request.Context(), emailInfo); err == nil {
		rsp.ID = id
	} else {
		log.Printf("CreateEmailTask fail, err:%v", err)
		return
	}
	c.JSON(200, Result{
		Data:       rsp,
		ResultCode: 200,
		Message:    "ok",
	})
}

// 查询名下的email
func queryEmail(c *gin.Context) {
	req := &model.QueryEmailReq{}
	if err := c.ShouldBind(req); err != nil {
		log.Printf("bind req failed, err:%v", err)
		c.JSON(200, Result{
			Data:       nil,
			ResultCode: 500,
			Message:    "param error",
		})
		return
	}
	if req.Creator == "" {
		req.Creator = c.Query("creator")
	}

	fmt.Printf("show query email req:%v", req)
	rsp := &model.QueryEmailRsp{}

	if emailList, err := middleware.GetDao().QueryEmailTask(c.Request.Context(), req.Creator); err == nil {
		for _, info := range emailList {
			basic := &model.EmailBasic{
				Receiver:    info.Receiver,
				Subject:     info.Subject,
				Mode:        info.Mode,
				Public:      info.Public,
				TriggleDate: time.Unix(int64(info.TriggleTime), 0).Format("2006-01-02"),
			}
			if info.State == 0 {
				basic.State = "等待投递"
			} else if info.State == 1 {
				basic.State = "投递成功"
			} else if info.State == 2 {
				basic.State = "投递失败"
			} else {
				basic.State = "未知状态"
			}
			rsp.BasicList = append(rsp.BasicList, *basic)
		}
	} else {
		log.Printf("CreateEmailTask fail, err:%v", err)
	}

	// emailInfo to EmailBasic
	c.JSON(200, Result{
		Data:       rsp,
		ResultCode: 200,
		Message:    "ok",
	})
}
