package emailx

import (
	"errors"
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"strings"
	"time"
	"volunteer-team/backend/internal/infrastructure/configs"
	"volunteer-team/backend/internal/infrastructure/global"
	"volunteer-team/backend/internal/infrastructure/pkg/syncx"
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
func (ex *EmailX) SendLoginCode(to string, code string) (err error) {
	ex.store.StoreWithTTL(to, code, ex.ttl)
	subject := fmt.Sprintf("[%s]邮箱登录", ex.config.Subject)
	text := fmt.Sprintf("你正在进行邮箱登录，登录的验证码是：%s，十分钟内有效", code)
	return ex.sendEmail(to, subject, text)
}
func (ex *EmailX) SendResetPwdCode(to string, code string) (err error) {
	ex.store.StoreWithTTL(to, code, ex.ttl)
	subject := fmt.Sprintf("[%s]重置密码", ex.config.Subject)
	text := fmt.Sprintf("你正在进行账号密码重置，重置的验证码是：%s，十分钟内有效", code)
	return ex.sendEmail(to, subject, text)
}

func (ex *EmailX) SendBindEmail(to string, code string) (err error) {
	ex.store.StoreWithTTL(to, code, ex.ttl)
	subject := fmt.Sprintf("[%s]绑定邮箱", ex.config.Subject)
	text := fmt.Sprintf("你正在进行邮箱绑定，绑定的验证码是：%s，十分钟内有效", code)
	return ex.sendEmail(to, subject, text)
}
func (ex *EmailX) sendEmail(to, subject, text string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", ex.config.SendNickname, ex.config.SendEmail)
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)

	addr := fmt.Sprintf("%s:%d", ex.config.Domain, ex.config.Port)
	auth := smtp.PlainAuth("", ex.config.SendEmail, ex.config.AuthCode, ex.config.Domain)

	type result struct{ err error }
	ch := make(chan result, 1)

	go func() {
		ch <- result{err: e.Send(addr, auth)}
	}()

	select {
	case res := <-ch:
		if res.err != nil && !strings.Contains(res.err.Error(), "short response") {
			return res.err
		} else {
			return nil
		}
	case <-time.After(10 * time.Second):
		return SEND_OVER_TIME
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
