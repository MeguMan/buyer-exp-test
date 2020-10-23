package main

import (
	"github.com/MeguMan/buyer-exp-test/internal/app/apiserver"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main(){
	databaseURL, exists := os.LookupEnv("DATABASE_URL")

	if exists {
		if err := apiserver.Start(databaseURL); err != nil {
			log.Fatal(err)
		}
	}
/*	file, err := os.Open("m_test_page.html")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)
	var text string
	for {
		_, err := file.Read(data)
		if err == io.EOF{   // если конец файла
			break           // выходим из цикла
		}
		text += string(data)
	}

	re := regexp.MustCompile(`itemProp="price" content="(.*)"`)
	priceString := re.FindAllString(text, -1)
	price, err := intFromTag(priceString[0])
	if err != nil {
		panic(err)
	}
	fmt.Println(price)*/
}

/*func intFromTag(tag string) (int, error) {
	re := regexp.MustCompile(`\d+`)
	nn := re.FindAllString(tag, -1)
	var str string
	for _, n := range nn {
		str += n
	}

	return strconv.Atoi(str)
}*/