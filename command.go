package main

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

type Command string

const (
	CMDSet Command = "SET"
	CMDGet Command = "GET"
)

type Message struct {
	Cmd   Command
	Key   []byte
	Value []byte
	TTL   time.Duration
}

func parseCommand(byteCmd []byte) (*Message, error) {
	var (
		strCmd = string(byteCmd)
		parts  = strings.Split(strCmd, " ")
	)
	if len(parts) < 2 {
		return nil, errors.New("invalid command format")
	}

	msg := &Message{
		Cmd: Command(parts[0]),
		Key: []byte(parts[1]),
	}
	if msg.Cmd == CMDSet {

		if len(parts) < 4 {
			return nil, errors.New("invalid set format")
		}
		msg.Value = []byte(parts[2])

		ttl, err := strconv.Atoi(parts[3])
		if err != nil {
			log.Println("invalid set ttl")
			return nil, err
		}
		msg.TTL = time.Duration(ttl)

	}

	return msg, nil
}
