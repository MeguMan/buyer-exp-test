package model

import (
	"fmt"
	"net/smtp"
	"strconv"
)

type User struct {
	ID    int
	Email string
}

func (u *User) SendEmail(ad *Ad) error {
	fmt.Println("ОТПРАВЛЯЮ ИМЕЙЛ")
	to := u.Email
	from := "dimon200019@gmail.com"
	pass := "*****"
	msg := to +
		"Subject: Уведомление о цене подписки\r\n" +
		"Цена на объявление " + ad.Link +
		" на данный момент составляет " + strconv.Itoa(ad.Price) + "\r\n" +
		"При изменении цены вам будет отправлено новое письмо"
	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")
	return smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))
}
