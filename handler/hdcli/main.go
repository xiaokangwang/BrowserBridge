package main

import (
	"fmt"
	"github.com/xiaokangwang/BrowserBridge/handler"
)

func main() {
	fmt.Println("V3")
	handler.Handle(&handler.HandleSettings{
		ListenAddr: "127.0.0.1:23321",
		RemoteAddr: "wss://ascendingdonky.herokuapp.com/r588qe7mu61e0jbqfd9uua5yqqsi91rrj6w2",
	})
}
