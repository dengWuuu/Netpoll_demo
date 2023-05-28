package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/netpoll"
)

func main() {
	network, address := "tcp", ":8080"
	listener, err := netpoll.CreateListener(network, address)
	if err != nil {
		panic("create net listener failed")
	}
	// init event loop
	eventLoop, _ := netpoll.NewEventLoop(
		handle,
		netpoll.WithOnPrepare(prepare),
		netpoll.WithReadTimeout(time.Second),
	)

	// start listen loop ...
	err = eventLoop.Serve(listener)
	if err != nil {
		panic("server start err")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = eventLoop.Shutdown(ctx)
	if err != nil {
		panic("shutdown err")
	}
}

var _ netpoll.OnPrepare = prepare
var _ netpoll.OnRequest = handle

func prepare(netpoll.Connection) context.Context {
	return context.Background()
}

func handle(ctx context.Context, connection netpoll.Connection) error {
	fmt.Printf("ctx: %v\n", ctx)
	reader := connection.Reader()
	defer func(reader netpoll.Reader) {
		err := reader.Release()
		if err != nil {
			panic("reader release err")
		}
	}(reader)
	msg, _ := reader.ReadString(reader.Len())
	println(msg)
	return nil
}
