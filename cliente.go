package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	nick     string
	sala     *sala
	comandos chan<- comando
}

func (c *client) lerInput() {
	for {
		mens, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		mens = strings.Trim(mens, "\r\n")
		args := strings.Split(mens, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/apelido":
			c.comandos <- comando{
				id:      CMD_NICK,
				cliente: c,
				args:    args,
			}
		case "/unir":
			c.comandos <- comando{
				id:      CMD_UNIR,
				cliente: c,
				args:    args,
			}
		case "/salas":
			c.comandos <- comando{
				id:      CMD_SALAS,
				cliente: c,
				args:    args,
			}
		case "/msg":
			c.comandos <- comando{
				id:      CMD_MSG,
				cliente: c,
				args:    args,
			}
		case "/sair":
			c.comandos <- comando{
				id:      CMD_SAIR,
				cliente: c,
			}
		default:
			c.err(fmt.Errorf("erro desconhecido: %s", cmd))
		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("ERRO: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}
