package client

import (
	"context"
	"go-cache/protocol"
	"log"
	"net"
)

type Options struct{}

type Client struct {
	conn net.Conn
}

func New(addr string, opts Options) (*Client, error) {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) Close() error {
	log.Println("closing client")
	return c.conn.Close()
}

func (c *Client) Set(ctx context.Context, key, value []byte, ttl int) (any, error) {

	cmd := &protocol.CommandSet{
		Key:   key,
		Value: value,
		TTL:   ttl,
	}

	_, err := c.conn.Write(cmd.Bytes())

	if err != nil {
		return nil, err
	}

	return nil, nil
}
