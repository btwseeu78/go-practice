package subroutine

import (
	"fmt"
	"net/http"
)

func TestLinks() {
	links := []string{
		"http://google.com",
		"http://golang.org",
	}

	for _, link := range links {
		go CheckLink(link)
	}
}

func CheckLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Printf("Calling site failed:%s with error %g/n", link, err)
		return
	}
	fmt.Printf("Sucessflly called : %s\n", link)
}
