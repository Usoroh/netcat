package main

import (
	"bufio"
	"net"
	"strings"
)

type room struct {
	name    string
	members map[net.Addr]*client
}

type client struct {
	conn     net.Conn
	username string
	room     *room
	commands chan<- command
}

func (c *client) msg(msg string) {
	c.conn.Write([]byte(">" + msg + "\n"))
}

func (r *room) broadcast(sender *client, msg string) {
	for addr, m := range r.members {
		if sender.conn.RemoteAddr() != addr {
			m.msg(msg)
		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("ERR: " + err.Error() + "\n"))
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/username":
			c.commands <- command{
				id:     "username",
				client: c,
				args:   args,
			}
		default:
			c.room.broadcast(c, cmd)
		}
	}
}
