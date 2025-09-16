package emailx

import (
	"context"
	"errors"
	"fmt"
	"net/smtp"
	"strings"
	"time"
	"volunteer-team/backend/internal/infrastructure/configs"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/pkg/syncx"

	"github.com/jordan-wright/email"
)

var (
	SEND_OVER_TIME = errors.New("发送超时")
)

type EmailX struct {
	config configs.Email
	ttl    time.Duration
	store  *syncx.Map[string, string]
}

func NewEmailX() *EmailX {
	return &EmailX{
		config: configs.Conf.Email,
		ttl:    10 * time.Minute,
		store:  global.CodeStore,
	}
}
func (ex *EmailX) SendLoginCode(ctx context.Context, to string, code string) (err error) {
	ex.store.StoreWithTTL(to, code, ex.ttl)
	subject := fmt.Sprintf("[%s]邮箱登录", ex.config.Subject)
	text := fmt.Sprintf("你正在进行邮箱登录，登录的验证码是：%s，十分钟内有效", code)
	return ex.sendEmail(ctx, to, subject, text)
}
func (ex *EmailX) SendResetPwdCode(ctx context.Context, to string, code string) (err error) {
	ex.store.StoreWithTTL(to, code, ex.ttl)
	subject := fmt.Sprintf("[%s]重置密码", ex.config.Subject)
	text := fmt.Sprintf("你正在进行账号密码重置，重置的验证码是：%s，十分钟内有效", code)
	return ex.sendEmail(ctx, to, subject, text)
}

func (ex *EmailX) SendBindEmail(ctx context.Context, to string, code string) (err error) {
	ex.store.StoreWithTTL(to, code, ex.ttl)
	subject := fmt.Sprintf("[%s]绑定邮箱", ex.config.Subject)
	text := fmt.Sprintf("你正在进行邮箱绑定，绑定的验证码是：%s，十分钟内有效", code)
	return ex.sendEmail(ctx, to, subject, text)
}

func (ex *EmailX) sendEmail(ctx context.Context, to, subject, text string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", ex.config.SendNickname, ex.config.SendEmail)
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)

	addr := fmt.Sprintf("%s:%d", ex.config.Domain, ex.config.Port)
	auth := smtp.PlainAuth("", ex.config.SendEmail, ex.config.AuthCode, ex.config.Domain)

	type result struct{ err error }
	done := make(chan result, 1)

	// 1. 计算剩余时间
	var timeout time.Duration
	if dl, ok := ctx.Deadline(); ok {
		timeout = time.Until(dl)
		if timeout <= 0 {
			return context.DeadlineExceeded
		}
	} else {
		timeout = 10 * time.Second // 调用方没给 deadline 就用默认
	}

	global.Log.Debug(timeout)

	// 2. 异步发送
	go func() {
		err := e.Send(addr, auth)
		// 过滤掉某些老旧服务器返回的“short response”伪错误
		if err != nil && !strings.Contains(err.Error(), "short response") {
			done <- result{err: err}
			return
		}
		done <- result{err: nil}
	}()

	// 3. 等待完成或超时 / 被取消
	select {
	case res := <-done:
		return res.err
	case <-time.After(timeout):
		return SEND_OVER_TIME
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (ex *EmailX) VerifyCode(email, code string) bool {
	var ans bool
	if v, ok := ex.store.Load(email); ok {
		if v == code {
			ans = true
			ex.store.Delete(email)
		}
	}
	return ans
}
