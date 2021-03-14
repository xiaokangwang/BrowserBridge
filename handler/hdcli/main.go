package main

import (
	"fmt"
	"github.com/xiaokangwang/BrowserBridge/handler"
)

func main() {
	fmt.Println("V3")
	handler.Handle(&handler.HandleSettings{
		ListenAddr: "127.0.0.1:23321",
		RemoteAddr: "wss://127.0.0.1:0",
	})
}
