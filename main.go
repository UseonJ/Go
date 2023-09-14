package main

import (
	"fmt"
	"time"
)

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

	c := make(chan bool)
	people := []string{"Ace", "Bear", "Chache", "Dag"}
	for _, person := range people {
		go isSexy(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}

}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 2)
	fmt.Println(person)
	c <- true
}
