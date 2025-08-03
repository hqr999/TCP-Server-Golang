package main

import (
	"log"
	"net"
)

type servidor struct {
	salas    map[string]*sala
	comandos chan comando
}

func novoServidor() *servidor {
	return &servidor{
		salas:    make(map[string]*sala),
		comandos: make(chan comando),
	}
}

func (s servidor) novoCliente(conn net.Conn){
		log.Printf("Novo cliente foi conectado: %s",conn.RemoteAddr().String())
		
	c := client{
		conn: conn,
		nick: "Anonimo",
		comandos: s.comandos,
			
	}

} 
