package subroutine

import (
	"fmt"
	"net/http"
)

func TestLinks() {
	links := []string{
		"https://google.com",
		"https://golang.org",
	}
	c := make(chan string)
	for _, link := range links {
		go CheckLink(link, c)
	}

	for i := range c {
		fmt.Println(i)
	}
}

func CheckLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Printf("Calling site failed:%s with error %g/n", link, err)
		c <- "The Url Might Be Down"
		return
	}
	fmt.Printf("Sucessflly called : %s\n", link)
	c <- "The uRl is working just fine"
}
