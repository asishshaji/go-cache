package main

import (
	"context"
	"fmt"
	"go-cache/cache"
	"log"
	"net"
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

	defer func() {
		conn.Close()
	}()

	buf := make([]byte, 2048)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("reading failure: %s", err)
			break
		}

		msg := buf[:n]
		fmt.Println(string(msg))
		go s.handleCmd(conn, buf[:n])
	}
}

func (s *Server) handleCmd(conn net.Conn, byteCmd []byte) {
	msg, err := parseCommand(byteCmd)
	if err != nil {
		log.Println("failed to parse command", err)
		return
	}

	switch msg.Cmd {
	case CMDSet:
		err = s.handleSetCmd(conn, msg)
	case CMDGet:
		err = s.handleGetCmd(conn, msg)
	}
	if err != nil {
		log.Println("failed to handle command:", err.Error())
		return
	}

}

func (s *Server) handleGetCmd(conn net.Conn, msg *Message) error {
	buf, err := s.cache.Get(msg.Key)
	if err != nil {
		return err
	}
	_, err = conn.Write(buf)
	return err
}

func (s *Server) handleSetCmd(conn net.Conn, msg *Message) error {
	if err := s.cache.Set(msg.Key, msg.Value, msg.TTL); err != nil {
		return err
	}

	go s.sendToFollowers(context.TODO(), msg)

	return nil
}

func (s *Server) sendToFollowers(ctx context.Context, msg *Message) error {
	return nil
}
