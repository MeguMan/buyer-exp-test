package emailsender

import "testing"

func TestSender(t *testing.T) *Sender {
	t.Helper()

	return &Sender{
		Email:    "buyerexptest@gmail.com",
		Password: "135798642qq",
		TLSPort:  "587",
	}
}
