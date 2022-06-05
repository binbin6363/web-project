package main

import (
	"fmt"
	"log"
	"net/smtp"
	"time"

	"gopkg.in/gomail.v2"
)

func smtpSend() {
	addr := "smtp.yeah.net:25"
	auth := smtp.PlainAuth("", "tpo2022@yeah.net", "AXNJPCHCMCMVVKDW", "smtp.yeah.net")
	from := "tpo2022@yeah.net"
	to := []string{"136800719@qq.com"}
	subject := fmt.Sprintf("Subject: %s\r\n", "核对服务器故障整改信息")
	send := fmt.Sprintf("From: %s 时光邮局\r\n", "tpo2022@yeah.net")
	receiver := fmt.Sprintf("To: %s\r\n", "136800719@qq.com")
	content_type := "Content-Type: text/html" + "; charset=UTF-8\r\n\r\n"
	content := "嗨你好，最近新建了签名，包含图片，请查收确认效果。"
	msg := []byte(subject + send + receiver + content_type + content)
	err := smtp.SendMail(addr, auth, from, to, msg)
	fmt.Println(err)
}

func gomailSend() {
	m := gomail.NewMessage()

	//发送人
	m.SetHeader("From", m.FormatAddress("tpo2022@yeah.net", "时光邮局"))
	//接收人
	m.SetHeader("To", "136800719@qq.com")
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", "核对服务器故障整改信息")
	//内容
	m.SetBody("text/html", "<h1>新年快乐。嗨你好，最近新建了签名，包含图片，请查收确认效果。</h1>")
	//附件
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer("smtp.yeah.net", 25, "tpo2022@yeah.net", "AXNJPCHCMCMVVKDW")

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		panic(err)
	}
	fmt.Printf("send mail success\n")

}

func main() {

	the_time, _ := time.ParseInLocation("2006-01-02T15:04:05.000Z", "2022-06-05T16:00:00.000Z", time.Local)
	log.Printf("time:%v", the_time.Unix())
	//gomailSend()
}
