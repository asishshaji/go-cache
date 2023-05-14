package main

import (
	"flag"
	"go-cache/cache"
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

	server := NewServer(opts, cache.New())
	server.Start()
}
