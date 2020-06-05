package main

type command struct {
	id     string
	client *client
	args   []string
}
