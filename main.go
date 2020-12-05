package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()//Cria o server
	go s.run()//manda o server rodar as funções da lista de comandos
//Faz o server esperar pelo Client
	colorRed := "\033[31m"
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
	defer listener.Close()
	//log.Printf("server started on :8888")
    log.Println(string(colorRed), "server started on :8888")
	
	for {
		//For para sempre que fica esperando o Cliente conectar
		conn, err := listener.Accept()//Cliente connectando
		if err != nil {
			//log.Printf("failed to accept connection: %s", err.Error())
			log.Println(string(colorRed), "failed to accept connection: %s", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}
