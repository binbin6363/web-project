package middleware

import (
	"log"

	"github.com/binbin6363/web-project/server/config"
	"github.com/binbin6363/web-project/server/model"
	"gopkg.in/gomail.v2"
)

// EmailSend 发送email
func EmailSend(info *model.EmailInfo) error {
	if info.Receiver == "" || info.Body == "" {
		log.Printf("Receiver|Body empty, no send yet, wait for implement")
		return nil
	}
	m := gomail.NewMessage()

	//发送人
	m.SetHeader("From", m.FormatAddress(config.GetConfig().Email.From, config.GetConfig().Email.NickName))
	//接收人
	m.SetHeader("To", info.Receiver)
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", info.Subject)
	//内容
	m.SetBody("text/html", info.Body)
	//附件
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer(config.GetConfig().Email.SmtpHost, config.GetConfig().Email.SmtpPort,
		config.GetConfig().Email.From, config.GetConfig().Email.Token)

	// 发送邮件
	err := d.DialAndSend(m)
	if err != nil {
		log.Printf("DialAndSend err %v:", err)
		return err
	}
	log.Printf("send mail success\n")
	return nil
}
