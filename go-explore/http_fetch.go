package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main()  {
	resp, err := http.Get("https://www.baidu.com")
	checkError(err)
	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	fmt.Println("result:",string(bytes))
}

func checkError(err error)  {
	if err!=nil{
		log.Fatalln("err:%v",err)
	}
}