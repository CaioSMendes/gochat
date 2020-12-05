package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()//Cria o server
	go s.run()//manda o server rodar as funções da lista de comandos
//Faz o server esperar pelo Client
	listener, err := net.Listen("tcp", ":8888")
	colorRed := "\033[31m"
	if err != nil {
		log.Fatalf(string(colorRed), "unable to start server: %s", err.Error())
	}
	defer listener.Close()
	//log.Printf("server started on :8888")
	log.Println(string(colorRed), "server started on :8888")
	for {
		//For para sempre que fica esperando o Cliente conectar
		conn, err := listener.Accept()//Cliente connectando
		if err != nil {
			log.Printf(string(colorRed), "failed to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
