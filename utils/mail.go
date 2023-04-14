package utils

import (
	"crypto/tls"
	"fmt"

	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"gopkg.in/gomail.v2"
)

func Send(mailTo, subject, message string) (success bool) {
	host, _ := config.String("mail::host")
	port, _ := config.Int("mail::port")
	userName, _ := config.String("mail::username")
	password, _ := config.String("mail::password")

	m := gomail.NewMessage()
	m.SetHeader("From", userName) // 发件人
	// m.SetHeader("From", "alias"+"<"+sendName+">") // 增加发件人别名

	m.SetHeader("To", mailTo) // 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
	// m.SetHeader("Cc", "******")                  // 抄送，可以多个
	// m.SetHeader("Bcc", "******")                 // 暗送，可以多个
	m.SetHeader("Subject", subject) // 邮件主题

	// text/html 的意思是将文件的 content-type 设置为 text/html 的形式，浏览器在获取到这种文件时会自动调用html的解析器对文件进行相应的处理。
	// 可以通过 text/html 处理文本格式进行特殊处理，如换行、缩进、加粗等等
	m.SetBody("text/html", message)

	// text/plain的意思是将文件设置为纯文本的形式，浏览器在获取到这种文件时并不会对其进行处理
	// m.SetBody("text/plain", "纯文本")
	// m.Attach("test.sh")   // 附件文件，可以是文件，照片，视频等等
	// m.Attach("lolcatVideo.mp4") // 视频
	// m.Attach("lolcat.jpg") // 照片

	d := gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)

	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		errMsg := fmt.Errorf("send mail failed, %w", err)
		logs.Info(errMsg)
		success = false
	} else {
		success = true
	}

	return success
}
