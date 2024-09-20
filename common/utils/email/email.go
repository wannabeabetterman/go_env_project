/*
*

	@author Yeoman
	@date:2022/4/12
	@Description
*/
package email

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

type Config struct {
	To       string
	Port     int
	From     string
	Host     string
	IsSSL    bool
	Secret   string
	Nickname string
	Address  string
}

// Email
// @Description:无附件邮箱发送
// @param c
// @param subject
// @param body
// @return error
func Email(c Config, subject string, body string) error {
	to := strings.Split(c.To, ",")
	return send(c, to, subject, body)
}

// EmailFile
// @Description: 有附件邮箱发送
// @param subject
// @param body
// @param r
// @param filename
// @return error
func EmailFile(c Config, subject string, body string, r io.Reader, filename string) error {
	to := strings.Split(c.To, ",")
	return sendFile(c, to, subject, body, r, filename)
}

// ErrorToEmail
// @Description: 给email中间件错误发送邮件到指定邮箱
// @param c
// @param subject
// @param body
// @return error
func ErrorToEmail(c Config, subject string, body string) error {
	to := strings.Split(c.To, ",")
	if to[len(to)-1] == "" { // 判断切片的最后一个元素是否为空,为空则移除
		to = to[:len(to)-1]
	}
	return send(c, to, subject, body)
}

// sendFile
// @Description: 有附件邮箱发送
// @param c
// @param to
// @param subject
// @param body
// @param r
// @param filename
// @return error
func sendFile(c Config, to []string, subject string, body string, r io.Reader, filename string) error {
	auth := smtp.PlainAuth("", c.From, c.Secret, c.Host)
	e := email.NewEmail()
	if c.Nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", c.Nickname, c.From)
	} else {
		e.From = c.From
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)

	e.Attach(r, filename, "text/plain; charset=utf-8")
	//e.Attachments=attachment
	var err error
	hostAddr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	if c.IsSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: c.Host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}

// send
// @Description: 无附件发送
// @param c
// @param to
// @param subject
// @param body
// @return error
func send(c Config, to []string, subject string, body string) error {
	auth := smtp.PlainAuth("", c.From, c.Secret, c.Host)
	e := email.NewEmail()
	if c.Nickname != "" {
		e.From = fmt.Sprintf("%s <%s>", c.Nickname, c.From)
	} else {
		e.From = c.From
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	if c.IsSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: c.Host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err
}

type ContentWriter struct {
	Contents []byte
}

func (cw *ContentWriter) String() string {
	if len(cw.Contents) > 0 {
		return string(cw.Contents)
	}
	return ""
}
func (cw *ContentWriter) Write(p []byte) (n int, err error) {
	if len(p) > 0 {
		cw.Contents = append(cw.Contents, p...)
		n = len(p)
	}
	return
}
