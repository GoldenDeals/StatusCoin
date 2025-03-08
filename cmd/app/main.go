package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/GoldenDeals/StatusCoin/cmd"
	"github.com/GoldenDeals/StatusCoin/internal/node"
	"github.com/GoldenDeals/StatusCoin/internal/share"
	"github.com/GoldenDeals/StatusCoin/internal/share/shutdown"
)

func main() {
	cmd.Configure()
	ctx, cancel := context.WithCancel(context.Background())

	shut := shutdown.Init(ctx)
	log := share.NewLogger("main", shut)
	log.Info("Main has been entered!")

	nodeEnt := node.Init(shut)
	go nodeEnt.Start(ctx)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Warn("Recived signal... Shutting down")
	cancel()
}
