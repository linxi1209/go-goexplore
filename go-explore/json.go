package main

import (
	"encoding/json"
	"fmt"
)

//JSON，XML，gob，Google 缓冲协议等等。Go 语言支持所有这些编码格式

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "ChengDu", "China"}
	wa := &Address{"private", "beijing", "China"}
	vc := &VCard{"Max", "Zhong", []*Address{pa, wa}, "none"}

	bytes, _ := json.Marshal(vc)
	fmt.Printf("json: %s", bytes)
}
