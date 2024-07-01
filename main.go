package main

import (
	"fmt"
	"github.com/imroc/req"
)

func main() {
	url := "https://httpbin.org/get"
	r, err := req.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", r)
}
