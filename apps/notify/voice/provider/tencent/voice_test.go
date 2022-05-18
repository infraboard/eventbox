package tencent_test

import (
	"os"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"

	"github.com/infraboard/eventbox/apps/notify/voice/provider/tencent"
)

var (
	// 专门用于语言通知
	txSecretId, txSecretKey string
)

var (
	// 查询你的Appid https://console.cloud.tencent.com/vms,
	voiceSdkAppId = "1400121047"
	// 查询你的模版Id https://console.cloud.tencent.com/vms/template
	voiceTemplateId = "168501"
)

func TestQcloudVoice(t *testing.T) {
	should := assert.New(t)
	v, err := tencent.NewQcloudVoice(txSecretId, txSecretKey, voiceSdkAppId)
	if should.NoError(err) {
		req := tencent.NewPhoneCallRequest("+8618108053819", voiceTemplateId, []string{"测试"})
		_, err := v.PhoneCall(req)
		should.NoError(err)
	}
}

func init() {
	zap.DevelopmentSetup()
	txSecretId = os.Getenv("TX_SECRET_ID")
	txSecretKey = os.Getenv("TX_SECRET_KEY")
}
