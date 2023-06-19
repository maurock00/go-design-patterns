package main

import (
	"net/http"
)

func main() {
	urls := []string{"https://www.google.com", "https://www.facebook.com", "https://www.udemy.com"}

	c := make(chan string)

	for _, website := range urls {
		go checkUrl(website, c)
	}

	variable := <-c
	println(variable)

	//for msg := range c {
	//	go func(link string) {
	//		time.Sleep(2 * time.Second)
	//		checkUrl(link, c)
	//	}(msg)
	//}
}

func checkUrl(website string, c chan string) {
	_, httpError := http.Get(website)

	if httpError != nil {
		println(httpError.Error())
		c <- website
	}

	println(website + " It's OK")
	c <- website
}
