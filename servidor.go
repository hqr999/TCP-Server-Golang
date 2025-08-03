package main

import (
	"fmt"
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

func (s *servidor) rodar() {
	for cmd := range s.comandos {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.cliente, cmd.args)
		case CMD_MSG:
			s.msg(cmd.cliente, cmd.args)
		case CMD_UNIR:
			s.unir(cmd.cliente, cmd.args)
		case CMD_SALAS:
			s.lista_salas(cmd.cliente, cmd.args)
		case CMD_SAIR:
			s.sair(cmd.cliente, cmd.args)

		}
	}
}

func (s *servidor) novoCliente(conn net.Conn) {
	log.Printf("Novo cliente foi conectado: %s", conn.RemoteAddr().String())

	c := client{
		conn:     conn,
		nick:     "Anonimo",
		comandos: s.comandos,
	}
	c.lerInput()
}

func (s *servidor) nick(c *client, args []string) {
	c.nick = args[1]
	c.msg(fmt.Sprintf("Tudo certo, eu vou chamá-lo de %s", c.nick))
}
func (s *servidor) unir(c *client, args []string) {
	nomeSala := args[1]

	r, ok := s.salas[nomeSala]
	if !ok {
		r = &sala{
			nome:    nomeSala,
			membros: make(map[net.Addr]*client),
		}
		s.salas[nomeSala] = r
	}
	r.membros[c.conn.RemoteAddr()] = c 
	if c.sala != nil {
		
	}
	c.conn = r 
}
func (s *servidor) lista_salas(c *client, args []string) {}
func (s *servidor) msg(c *client, args []string)         {}
func (s *servidor) sair(c *client, args []string)        {}

func (s *servidor) sairSalaAtual(c *client) {
	if c.conn != nil {
		delete(c.sala.membros,c.conn.RemoteAddr())
		c.sala.broadcast(c,fmt.Sprintf("%s deixou a sala",c.nick))
	}
}
