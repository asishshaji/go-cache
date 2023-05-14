package main

import (
	"flag"
	"go-cache/cache"
	"log"
	"net"
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
		IsLeader:   listenAddr == leaderAddr,
		LeaderAddr: *leaderAddr,
	}

	go func() {
		time.Sleep(time.Second * 2)
		conn, err := net.Dial("tcp", ":3000")
		if err != nil {
			log.Fatal(err)
		}

		conn.Write([]byte("SET Foo Bar 20000000000"))

		time.Sleep(time.Second * 2)

		conn.Write([]byte("GET Foo"))
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		val := buf[:n]
		log.Println(string(val))
	}()

	server := NewServer(opts, cache.New())
	server.Start()
}
