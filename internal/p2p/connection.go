package p2p

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

func (n *Host) Connect(peerAddr string) {
	peerMA, err := multiaddr.NewMultiaddr(peerAddr)
	if err != nil {
		panic(err)
	}
	peerAddrInfo, err := peer.AddrInfoFromP2pAddr(peerMA)
	if err != nil {
		panic(err)
	}

	// Connect to the node at the given address.
	if err := n.H.Connect(context.Background(), *peerAddrInfo); err != nil {
		panic(err)
	}
	fmt.Println("Connected to", peerAddrInfo.String())

	// Open a stream with the given peer.
	s, err := n.H.NewStream(context.Background(), peerAddrInfo.ID, protocolID)
	if err != nil {
		panic(err)
	}

	// Start the write and read threads.
	go writeCounter(s)
	go readCounter(s)
}
