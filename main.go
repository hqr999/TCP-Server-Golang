package main

import (
	"log"
	"net"
)

func main() {
		s := novoServidor()

		listener, err := net.Listen("tcp",":8080")
		if err != nil {
				log.Fatal("Incapax=z de iniciar o servidor",err.Error())
	}

	defer listener.Close()
	log.Println("Servidor inicializado na porta :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Incapaz de aceitar conexões: %s",err.Error())
			continue
		}

		go s.novoCliente()
	}	

}
