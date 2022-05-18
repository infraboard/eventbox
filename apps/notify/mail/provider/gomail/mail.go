package gomail

import (
	"fmt"
	"strings"
	"sync"

	"github.com/go-gomail/gomail"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/eventbox/apps/notify/mail"
)

const (
	ERR_BROKEN_CONN = "write: broken pipe"
)

func NewSender(host string, port int, username, password string) *Sender {
	s := &Sender{
		host:     host,
		port:     port,
		username: username,
		password: password,
		log:      zap.L().Named("Mail Sender"),
	}
	return s
}

type Sender struct {
	host     string
	port     int
	username string
	password string

	log    logger.Logger
	sender gomail.SendCloser
	lock   sync.Mutex
}

func (s *Sender) init() error {
	d := gomail.NewDialer(s.host, s.port, s.username, s.password)
	sender, err := d.Dial()
	if err != nil {
		return err
	}
	s.sender = sender
	return nil
}

func (s *Sender) Send(req *mail.Request) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.sender == nil {
		if err := s.init(); err != nil {
			return err
		}
	}

	m := gomail.NewMessage()
	m.SetHeader("From", s.username)
	m.SetHeader("To", req.To...)
	m.SetHeader("Subject", req.Title)
	m.SetBody("text/html", req.Content)

	err := gomail.Send(s.sender, m)
	if err != nil {
		// 如果sender链接异常, 尝试重链
		if strings.Contains(err.Error(), ERR_BROKEN_CONN) {
			if err := s.init(); err != nil {
				return fmt.Errorf("Sender 初始化异常, %s", err)
			}

			return gomail.Send(s.sender, m)
		}

		return err
	}

	return nil
}
