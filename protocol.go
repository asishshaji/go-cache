package main

import (
	"errors"
	"fmt"
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

func (msg *Message) ToBytes() []byte {
	var cmd string
	switch msg.Cmd {
	case CMDSet:
		cmd = fmt.Sprintf("%s %s %s %d", msg.Cmd, msg.Key, msg.Value, msg.TTL)

	case CMDGet:
		cmd = fmt.Sprintf("%s %s", msg.Cmd, msg.Key)

	default:
		log.Panic("unknown command")
	}
	return []byte(cmd)
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
