package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"https://www.baidu.com",
	"https://blog.csdn.net/daybreak1209/article/details/51307897",
	"http://blog.golang.org/",
}

func main() {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.print("Error:", url, err)
		}
		fmt.Println("response:", resp.Status)
	}
}
