package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
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
			s.lista_salas(cmd.cliente)
		case CMD_SAIR:
			s.sair(cmd.cliente)

		}
	}
}

func (s *servidor) novoCliente(conn net.Conn) {
	log.Printf("Novo cliente foi conectado: %s", conn.RemoteAddr().String())

	c := &client{
		conn:     conn,
		nick:     "Anonimo",
		comandos: s.comandos,
	}

	c.lerInput()
}

func (s *servidor) nick(c *client, args []string) {
	c.nick = args[1]
	c.msg(fmt.Sprintf("Tudo certo, vou chamá-lo de %s", c.nick))
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


	
	if c.sala != nil {
		s.sairSalaAtual(c)
	}
	r.membros[c.conn.RemoteAddr()] = c
	c.sala = r
	r.broadcast(c, fmt.Sprintf("%s se uniu a sala", c.nick))
	c.msg(fmt.Sprintf("Bem Vindo %s", r.nome))
}

func (s *servidor) lista_salas(c *client) {
	var salas []string
	for nome := range s.salas {
		salas = append(salas, nome)
	}

	c.msg(fmt.Sprintf("Salas disponíveis são: %s", strings.Join(salas, ",")))

}



func (s *servidor) msg(c *client, args []string) {
	if c.sala == nil {
		c.err(errors.New("Você precisa se unir a uma sala primeiro"))
		return
	}
	c.sala.broadcast(c, c.nick+": "+strings.Join(args[1:], " "))

}


func (s *servidor) sair(c *client) {
	log.Printf("Cliente foi desconectado: %s", c.conn.RemoteAddr().String())
	s.sairSalaAtual(c)
	c.msg("Triste em saber que você foi embora =( ")
	c.conn.Close()

}

func (s *servidor) sairSalaAtual(c *client) {
	if c.conn != nil {
		delete(c.sala.membros, c.conn.RemoteAddr())
		c.sala.broadcast(c, fmt.Sprintf("%s deixou a sala", c.nick))
	}
}
