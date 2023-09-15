package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://kr.indeed.com/jobs?q=%ED%8C%8C%EC%9D%B4%EC%8D%AC&limit50=&vjk=d19b63e763f4cab0"

func main() {
	// account := accounts.NewAccount("nicolas")
	// account.Deposit(10)
	// fmt.Println(account.Balance())
	// err := account.Withdraw(20)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())
	// 	dictionary := mydict.Dictionary{}
	// 	baseWord := "hello"
	// 	dictionary.Add(baseWord, "First")
	// 	dictionary.Search(baseWord)
	// 	dictionary.Delete(baseWord)
	// 	word, err := dictionary.Search(baseWord)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(word)
	// 	}
	// }

	// 	c := make(chan bool)
	// 	people := []string{"Ace", "Bear", "Chache", "Dag"}
	// 	for _, person := range people {
	// 		go isSexy(person, c)
	// 	}

	// 	for i := 0; i < len(people); i++ {
	// 		fmt.Println(<-c)
	// 	}

	// }

	// func isSexy(person string, c chan bool) {
	// 	time.Sleep(time.Second * 2)
	// 	fmt.Println(person)
	// 	c <- true

	// #4 get Pages

	getPages()

}

func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	fmt.Println(doc)

	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
