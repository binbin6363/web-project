package middleware

import (
	"context"
	"log"
	"time"

	"github.com/binbin6363/web-project/server/model"
)

type TickTask struct {
	Name string
}

var tt *TickTask

func NewTick() *TickTask {
	tt = &TickTask{
		Name: "tick",
	}

	return tt
}

func (tt *TickTask) Tick() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Tick Recovered in f", r)
		}
	}()

	// 创建一个计时器
	timeTicker := time.Tick(time.Second * 10)
	for {
		<-timeTicker
		log.Printf("search %s task", tt.Name)
		taskList, err := GetDao().ScanEmailTask(context.Background())
		if err != nil {
			continue
		}

		for idx, task := range taskList {
			info := &model.EmailInfo{
				Receiver: task.Receiver,
				Subject:  task.Subject,
				Body:     task.Body,
			}
			if err := EmailSend(info); err != nil {
				log.Printf("[%d] send email err:%v", idx, err)
			} else {
				log.Printf("[%d] send email ok:%v", idx, info)
			}
		}
	}
}
