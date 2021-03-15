package main

import (
	"fmt"
	"github.com/xiaokangwang/BrowserBridge/handler"
	"os"
)

func main() {

	fmt.Println("V3")
	handler.Handle(&handler.HandleSettings{
		ListenAddr: os.Args[2],
		RemoteAddr: os.Args[1],
	})
}
