package model_test

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"regexp"
	"testing"
)

func TestAd_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		a       func() *model.Ad
		isValid bool
	}{
		{
			name: "valid",
			a: func() *model.Ad {
				return model.TestAd(t)
			},
			isValid: true,
		},
		{
			name: "invalid link",
			a: func() *model.Ad {
				a := model.TestAd(t)
				a.Link = "invalidLink"
				return a
			},
			isValid: false,
		},
		{
			name: "empty link",
			a: func() *model.Ad {
				a := model.TestAd(t)
				a.Link = ""
				return a
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.a().Validate())
			} else {
				assert.Error(t, tc.a().Validate())
			}
		})
	}
}

func TestAd_ParsePrice(t *testing.T) {
	file, err := ioutil.ReadFile("avito_page.html")
	if err != nil {
		log.Println(err)
	}
	text := string(file)

	re := regexp.MustCompile(`itemProp="price" content="(.*)"`)
	priceString := re.FindAllString(text, -1)
	price, err := model.IntFromTag(priceString[0])
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 1199000, price)
}
