package bridge

import (
	"fmt"
	"github.com/gopherjs/websocket"
	"github.com/xiaokangwang/BrowserBridge/proto"
	"github.com/xtaci/smux"
	"io"
	"time"
)

type Settings struct {
	DialAddr string
}

func Bridge(s *Settings) {
	for {
		delay := time.NewTimer(time.Second)
		DoConnect := func() {
			conn, err := websocket.Dial(s.DialAddr)
			if err != nil {
				return
			}
			smuxc, err := smux.Client(conn, nil)
			if err != nil {
				return
			}
			for {
				stream, err := smuxc.Accept()
				if err != nil {
					fmt.Println(err)
					return
				}
				go func() {
					err, req := proto.ReadRequest(stream)
					if err != nil {
						fmt.Println(err)
						return
					}
					conn2, err := websocket.Dial(req.Destination)
					if err != nil {
						fmt.Println(err)
						stream.Close()
						return
					}

					go io.Copy(stream, conn2)
					io.Copy(conn2, stream)
					stream.Close()
				}()

			}
		}
		DoConnect()
		<-delay.C
	}

}
