package main

import (
	"context"
	"flag"
	"go-cache/cache"
	"go-cache/client"
	"log"

	"time"
)

func main() {

	var (
		listenAddr = flag.String("listenaddr", ":3000", "listen port of server")
		leaderAddr = flag.String("leaderaddr", "", "listen port of the leader")
	)
	flag.Parse()

	opts := ServerOpts{
		ListenAddr: *listenAddr,
		IsLeader:   len(*leaderAddr) == 0,
		LeaderAddr: *leaderAddr,
	}

	go func() {
		time.Sleep(time.Second * 3)
		for i := 0; i < 12; i++ {
			SendCommand()
			time.Sleep(time.Millisecond * 200)
		}
	}()

	server := NewServer(opts, cache.New())
	server.Start()
}

func SendCommand() {
	client, _ := client.New(":3000", client.Options{})

	_, err := client.Set(context.Background(), []byte("foo"), []byte("bar"), 2)
	if err != nil {
		log.Fatal(err)
	}

	client.Close()
}
