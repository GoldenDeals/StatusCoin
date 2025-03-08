package p2p

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
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

		err := binary.Write(s, binary.BigEndian, counter)
		if err != nil {
			panic(err)
		}
	}
}

func readCounter(s network.Stream) {
	for {
		var counter uint64

		err := binary.Read(s, binary.BigEndian, &counter)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Received %d from %s\n", counter, s.ID())
	}
}
