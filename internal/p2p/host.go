package p2p

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

const protocolID = protocol.ID("/protochat/1.0.0")

type Host struct {
	H    host.Host
	MDNS *mdns.Service
}

func Init() *Host {
	node, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"))
	if err != nil {
		panic(err)
	}
	// print the node's PeerInfo in multiaddr format
	peerInfo := peer.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}
	addrs, err := peer.AddrInfoToP2pAddrs(&peerInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("libp2p node address:", addrs[0])

	node.SetStreamHandler(protocolID, handleStream)

	return &Host{H: node}
}

func (n *Host) Shutdown(_ string) error {
	if err := n.H.Close(); err != nil {
		return err
	}

	return nil
}
