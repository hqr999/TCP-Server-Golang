package main

import "net"

type sala struct {
		nome string 
		membros map[net.Addr]*client
}
