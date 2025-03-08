package node

import (
	"context"
	"os"

	"github.com/GoldenDeals/StatusCoin/internal/p2p"
	"github.com/GoldenDeals/StatusCoin/internal/share/shutdown"
)

type Node struct {
	SD *shutdown.Shutdown
}

func Init(sd *shutdown.Shutdown) *Node {
	return &Node{SD: sd}
}

func (n *Node) Start(ctx context.Context) {
	host := p2p.Init()
	if len(os.Args) > 1 {
		host.Connect(os.Args[1])
	}

	n.SD.Push("node-p2p", host.Shutdown)
	<-ctx.Done()
}
