package main 

type comandoID int 

const (
	CMD_NICK comandoID = iota
	CMD_UNIR
	CMD_SALAS
	CMD_MSG
	CMD_SAIR
)

type comando struct {
		id comandoID
		cliente *cliente 
		args []string
}
