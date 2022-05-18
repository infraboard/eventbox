package gomail_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/infraboard/eventbox/apps/notify/mail"
	"github.com/infraboard/eventbox/apps/notify/mail/provider/gomail"
)

var sender *gomail.Sender

func TestSend(t *testing.T) {
	m := &mail.Request{
		To:      []string{"yumaojun@g7.com.cn"},
		Title:   "告警告警",
		Content: "告警来了",
	}

	err := sender.Send(m)
	if err != nil {
		panic(err)
	}
}

func init() {
	h := os.Getenv("MAIL_SMTP_HOST")
	p := os.Getenv("MAIL_SMTP_PORT")
	pi, _ := strconv.Atoi(p)
	user := os.Getenv("MAIL_USERNAME")
	pass := os.Getenv("MAIL_PASSWORD")
	sender = gomail.NewSender(h, pi, user, pass)
}
