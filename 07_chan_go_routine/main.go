package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://xvideos.com",
		"http://onlinesb.com",
		"http://golang.org",
	}
	c := make(chan string)

	for _, link := range links {
		go checklink(link, c)
	}
	for l := range c { //i := 0; i < len(links); i++
		//fmt.Println(<-c)
		// go checklink(l, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checklink(link, c)
		}(l)
	}

}

func checklink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("link is down", link)
		//c <- "link is down I guess"
		c <- link
	}
	fmt.Println(link, "is up..")
	// c <- "Yep up..."
	c <- link
}
