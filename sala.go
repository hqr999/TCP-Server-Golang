package main

import "net"

type sala struct {
		nome string 
		membros map[net.Addr]*client
}

func (s *sala) broadcast(ss *client, msg string){
		for ender,m := range s.membros {
				if ender != ss.conn.RemoteAddr(){
				m.msg(msg)
		}
	}
}
