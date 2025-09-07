package test

import (
	"errors"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
	"volunteer-team/backend/internal/infrastructure/configs"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/pkg/emailx"
	"volunteer-team/backend/internal/infrastructure/utils/code"
)

func TestEmail(t *testing.T) {
	configs.LoadConfig()
	global.Init()
	c := code.GenCode()
	mail := "2896151192@qq.com"
	t.Log("验证码:", c)

	ex := emailx.NewEmailX()

	// 1. 先记录时间
	start := time.Now()
	err := ex.SendLoginCode(mail, c)
	elapsed := time.Since(start)

	t.Logf("SendLoginCode 耗时: %v, 返回错误: %v", elapsed, err)

	// 2. 如果耗时≈10s且err==SEND_OVER_TIME，说明e.Send内部卡死
	if errors.Is(err, emailx.SEND_OVER_TIME) {
		t.Log("触发超时，但邮件可能已发出——去邮箱确认")
	}

	// 3. 验证码存储逻辑是否OK
	global.CodeStore.Range(func(email, code string) bool {
		t.Logf("email: %s, code: %s\n", email, code)
		return true // 返回 false 会提前终止遍历
	})
	assert.Equal(t, true, ex.VerifyCode(mail, c))
	global.CodeStore.Range(func(email, code string) bool {
		t.Logf("email: %s, code: %s\n", email, code)
		return true // 返回 false 会提前终止遍历
	})
}
