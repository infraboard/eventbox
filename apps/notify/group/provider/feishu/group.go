package feishu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewSender() *Sender {
	return &Sender{
		client: &http.Client{Timeout: 3 * time.Second},
		log:    zap.L().Named("Feishu Group"),
	}
}

type Sender struct {
	client *http.Client
	log    logger.Logger
}

type Response struct {
	Code int             `json:"code"`
	Data json.RawMessage `json:"data"`
	Msg  string          `json:"msg"`
}

func (s *Sender) Send(m *NotifyMessage) error {
	messageObj := NewCardMessage(m)
	body, err := json.Marshal(messageObj)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", m.RobotURL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// 发起请求
	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("send request error, %s", err)
	}
	defer resp.Body.Close()

	// 读取body
	bytesB, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response error, %s", err)
	}
	respString := string(bytesB)
	if (resp.StatusCode / 100) != 2 {
		return fmt.Errorf("status code[%d] is not 200, response %s", resp.StatusCode, respString)
	}

	fsResp := new(Response)
	err = json.Unmarshal(bytesB, fsResp)
	if err != nil {
		return err
	}
	if fsResp.Code != 0 {
		return fmt.Errorf("send feisu msg error, code: %d, msg: %s", fsResp.Code, fsResp.Msg)
	}

	return nil
}
