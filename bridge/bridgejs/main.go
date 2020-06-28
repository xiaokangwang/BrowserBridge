package main

import (
	"fmt"
	"github.com/xiaokangwang/BrowserBridge/bridge"
)

func main() {
	fmt.Println("V3")
	go func() {
		bridge.Bridge(&bridge.Settings{DialAddr: "ws://127.0.0.1:23321/link"})
	}()
}

//export GOPHERJS_GOROOT="$(go1.12.16 env GOROOT)"
