package main

import (
	"fmt"
	"go-cache/cache"
	"go-cache/protocol"
	"io"
	"log"
	"net"
	"time"
)

type ServerOpts struct {
	ListenAddr string
	IsLeader   bool
	LeaderAddr string
}

type Server struct {
	ServerOpts
	cache cache.Cacher
}

func NewServer(opts ServerOpts, c cache.Cacher) *Server {
	return &Server{
		ServerOpts: opts,
		cache:      c,
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return fmt.Errorf("listen error:%s", err.Error())
	}

	log.Printf("server starting on port [%s]\n", s.ListenAddr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("accept error: %s\n", err)
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		cmd, err := protocol.ParseCommand(conn)
		if err != nil {
			if err == io.EOF {
				break
			}
			//drop the connection
			break
		}

		go s.handleCmd(conn, cmd)
	}

}

func (s *Server) handleCmd(conn net.Conn, cmd any) {

	switch v := cmd.(type) {
	case *protocol.CommandSet:
		err := s.handleSetCommand(conn, v)
		if err != nil {
			log.Println(err)
		}
	case *protocol.CommandGet:

	}
}

func (s *Server) handleSetCommand(conn net.Conn, cmd *protocol.CommandSet) error {
	log.Printf("SET %s to %s\n", cmd.Key, cmd.Value)
	return s.cache.Set(cmd.Key, cmd.Value, time.Duration(cmd.TTL))

}
