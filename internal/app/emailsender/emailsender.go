package emailsender

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
	"net/smtp"
	"strconv"
)

type Sender struct {
	Email    string
	Password string
	TLSPort  string
}

func New(s *Sender) *Sender {
	return &Sender{
		Email:    s.Email,
		Password: s.Password,
		TLSPort:  s.TLSPort,
	}
}

func (s *Sender) SendEmail(u *model.User, a *model.Ad) error {
	to := u.Email
	from := s.Email
	pass := s.Password
	msg := to +
		"Subject: Уведомление о цене подписки\r\n" +
		"Цена на объявление " + a.Link +
		" на данный момент составляет " + strconv.Itoa(a.Price) + "\r\n" +
		"При изменении цены вам будет отправлено новое письмо"
	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")
	return smtp.SendMail("smtp.gmail.com:"+s.TLSPort, auth, from, []string{to}, []byte(msg))
}
