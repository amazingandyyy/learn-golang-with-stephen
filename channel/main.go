package main

import "net/http"
import "fmt"

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://anazon.com",
	}

	c := make(chan string)

	for _, l := range links {
		go check(l, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func check(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link + " might be down"
		return
	}
	c <- link + " is up"
}
