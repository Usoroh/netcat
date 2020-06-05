package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {

	for cmd := range s.commands {
		switch cmd.id {
		case "username":
			s.username(cmd.client, cmd.args)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("New client has connected %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		username: "anonymous",
		commands: s.commands,
	}

	c.readInput()
}

func (s *server) username(c *client, args []string) {
	c.username = args[1]
	c.msg(fmt.Sprintf("All right, I will call you %s", c.username))
}

func (s *server) msg(c *client, args []string) {
	if c.room == nil {
		c.err(errors.New("You must join the room first"))
		return
	}
	c.room.broadcast(c, c.username+": "+strings.Join(args[1:len(args)], " "))
}
