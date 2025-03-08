package p2p

import (
	"fmt"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

const discoveryNamespace = "example"

func (n *Host) SetupDiscovery() {
	// Setup peer discovery.
	discoveryService := mdns.NewMdnsService(
		n.H,
		discoveryNamespace,
		&discoveryNotifee{},
	)
	defer discoveryService.Close()
}

type discoveryNotifee struct {
	h host.Host
}

func (n *discoveryNotifee) HandlePeerFound(peerInfo peer.AddrInfo) {
	fmt.Println("found peer", peerInfo.String())
}
