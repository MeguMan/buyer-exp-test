package emailsender

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSender_SendEmail(t *testing.T) {

	testCases := []struct {
		name    string
		s       func() *Sender
		a       *model.Ad
		u       *model.User
		isValid bool
	}{
		{
			name: "valid",
			s: func() *Sender {
				return TestSender(t)
			},
			a:       model.TestAd(t),
			u:       model.TestUser(t),
			isValid: true,
		},
		{
			name: "incorrect email",
			s: func() *Sender {
				s := TestSender(t)
				s.Email = "incorrectEmail"
				return s
			},
			a:       model.TestAd(t),
			u:       model.TestUser(t),
			isValid: false,
		},
		{
			name: "incorrect password",
			s: func() *Sender {
				s := TestSender(t)
				s.Password = "incorrectPassword"
				return s
			},
			a:       model.TestAd(t),
			u:       model.TestUser(t),
			isValid: false,
		},
		{
			name: "incorrect TLS port",
			s: func() *Sender {
				s := TestSender(t)
				s.TLSPort = "incorrectTLS"
				return s
			},
			a:       model.TestAd(t),
			u:       model.TestUser(t),
			isValid: false,
		},
		{
			name: "empty email",
			s: func() *Sender {
				s := TestSender(t)
				s.Email = ""
				return s
			},
			a:       model.TestAd(t),
			u:       model.TestUser(t),
			isValid: false,
		},
		{
			name: "empty password",
			s: func() *Sender {
				s := TestSender(t)
				s.Password = ""
				return s
			},
			a:       model.TestAd(t),
			u:       model.TestUser(t),
			isValid: false,
		},
		{
			name: "empty TLS port",
			s: func() *Sender {
				s := TestSender(t)
				s.TLSPort = ""
				return s
			},
			a:       model.TestAd(t),
			u:       model.TestUser(t),
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.s().SendEmail(tc.u, tc.a))
			} else {
				assert.Error(t, tc.s().SendEmail(tc.u, tc.a))
			}
		})
	}
}
