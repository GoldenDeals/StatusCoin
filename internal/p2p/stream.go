package p2p

import (
	"fmt"
	"time"

	"github.com/GoldenDeals/StatusCoin/internal/gen"
	"github.com/libp2p/go-libp2p/core/network"
	"google.golang.org/protobuf/proto"
)

func handleStream(stream network.Stream) {
	go writeCounter(stream)
	go readCounter(stream)
}

func writeCounter(s network.Stream) {
	var counter uint64

	for {
		<-time.After(time.Second)
		counter++

		mes := &gen.MyMessage{Content: fmt.Sprintf("Hello %d", counter), Timestamp: time.Now().Unix()}

		buff, err := proto.Marshal(mes)
		if err != nil {
			panic(err)
		}

		_, err = s.Write(buff)
		if err != nil {
			panic(err)
		}
	}
}

func readCounter(s network.Stream) {
	for {
		var mes gen.WrapperMessage

		buf := make([]byte, 1024)
		n, err := s.Read(buf)
		if err != nil {
			panic(err)
		}

		err = proto.Unmarshal(buf[:n], &mes)
		if err != nil {
			panic(err)
		}

		// fmt.Printf("Received %s:%d from %s\n", mes.Content, mes.Timestamp, s.ID())
	}
}
