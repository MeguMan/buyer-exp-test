package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"regexp"
	"strconv"
)

type Ad struct {
	ID    int
	Link  string
	Price int
}

func (a *Ad) Validate() error {
	return validation.ValidateStruct(a, validation.Field(&a.Link, validation.Required, is.URL))
}

func (a *Ad) ParsePrice(url string) int {
	//if strings.Contains(url, "www.") {
	//	url = strings.Replace(url, "www.", "m.", -1)
	//}
	//
	//req, err := http.Get(url)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//body, err := ioutil.ReadAll(req.Body)
	//
	//re := regexp.MustCompile(`itemProp="price" content="(.*)"`)
	//priceString := re.FindAllString(string(body), -1)
	//price, err := intFromTag(priceString[0])
	//if err != nil {
	//	panic(err)
	//}

	return 123
}

func IntFromTag(tag string) (int, error) {
	re := regexp.MustCompile(`\d+`)
	nn := re.FindAllString(tag, -1)
	var str string
	for _, n := range nn {
		str += n
	}

	return strconv.Atoi(str)
}
