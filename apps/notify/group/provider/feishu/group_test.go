package feishu_test

import (
	"testing"

	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/eventbox/apps/notify/group/provider/feishu"
)

var (
	sender *feishu.Sender
)

func TestSend(t *testing.T) {
	g1 := &feishu.FiledGroup{
		Items: []*feishu.NotifyFiled{
			{Key: "📅 时间: ", Value: "2022-02-23 20:17:51"},
			{Key: "👤 用户: ", Value: "测试"},
			{Key: "🔹 主机: ", Value: "host01"},
			{Key: "📋 操作: ", Value: "command xxx"},
		},
		EndType: feishu.FiledGroupEndType_Hr,
	}

	g2 := &feishu.FiledGroup{
		Items: []*feishu.NotifyFiled{
			{Key: "📅 时间: ", Value: "2022-02-23 20:17:51"},
			{Key: "👤 用户: ", Value: "测试"},
			{Key: "🔹 主机: ", Value: "host01"},
			{Key: "📋 操作: ", Value: "command xxx"},
		},
		EndType: feishu.FiledGroupEndType_Line,
	}

	message := feishu.NewFiledMarkdownMessage(
		"feishu robot addr",
		"1 级报警 - 数据平台",
		feishu.COLOR_RED,
		g1,
		g2,
	)
	message.Note = []string{"测试备注模块"}
	sender.Send(message)
}

func init() {
	zap.DevelopmentSetup()
	sender = feishu.NewSender()
}
