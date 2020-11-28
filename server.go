package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	members map[net.Addr]*client
	commands chan command
}	

func newServer() *server {
	return &server{
		// rooms:    make(map[string]*room),
		members : make(map[net.Addr]*client),
		commands: make(chan command),
	}
}
func (s *server) broadcast(sender *client, msg string) {//Modificar essa funçao para que ela funcione sem o room
	for addr, m := range s.members {
		if sender.conn.RemoteAddr() != addr {
			m.msg(msg)
		}
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NICK:
			s.nick(cmd.client, cmd.args[1])
		// case CMD_JOIN:
		// 	s.join(cmd.client, cmd.args[1])
		// case CMD_ROOMS:
		// 	s.listRooms(cmd.client)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT://Modificar
			s.quit(cmd.client)
		}
	}
}

func (s *server) newClient(conn net.Conn) {
	log.Printf("new client has joined: %s", conn.RemoteAddr().String())
	//Criar funçao "qual o seu nick" e setar variavel temporario para input do usuario
	c := &client{
		conn:     conn,
		nick:     "anonymous",//variavel temporario
		commands: s.commands,
	}
	s.members[c.conn.RemoteAddr()] = c
	//VERIFICAR IMPLEMENTAÇÃO DA BIBLIOTECA DE IO
	
	c.msg(fmt.Sprintf("Commands\n /msg to send messages\n /nick to change your nick;\n"))
	c.readInput()
}
//Essa funçao nick ja cria o nick
func (s *server) nick(c *client, nick string) {
	//adicionar printf: qual o seu nick
	c.nick = nick
	c.msg(fmt.Sprintf("all right, I will call you %s", nick))
}

// func (s *server) join(c *client, roomName string) {
// 	r, ok := s.rooms[roomName]
// 	if !ok {
// 		r = &room{
// 			name:    roomName,
// 			members: make(map[net.Addr]*client),
// 		}
// 		s.rooms[roomName] = r
// 	}
// 	r.members[c.conn.RemoteAddr()] = c

// 	s.quitCurrentRoom(c)
// 	c.room = r

// 	r.broadcast(c, fmt.Sprintf("%s joined the room", c.nick))

// 	c.msg(fmt.Sprintf("welcome to %s", roomName))
// }

// func (s *server) listRooms(c *client) {
// 	var rooms []string
// 	for name := range s.rooms {
// 		rooms = append(rooms, name)
// 	}

// 	c.msg(fmt.Sprintf("available rooms: %s", strings.Join(rooms, ", ")))
// }

func (s *server) msg(c *client, args []string) {
	msg := strings.Join(args[1:len(args)], " ")//une a mesnagem
	s.broadcast(c, c.nick+": "+msg)
	//s.broadcast(c, c.nick+": "+msg)
}

func (s *server) quit(c *client) {
	log.Printf("client has left the chat: %s", c.conn.RemoteAddr().String())
	//Sair do servidor
	// s.quitCurrentRoom(c)

	c.msg("sad to see you go =(")
	// c.conn.Close()
	c.conn.Close();
}

// func (s *server) quitCurrentRoom(c *client) {
// 	if c.room != nil {
// 		oldRoom := s.rooms[c.room.name]
// 		delete(s.rooms[c.room.name].members, c.conn.RemoteAddr())
// 		oldRoom.broadcast(c, fmt.Sprintf("%s has left the room", c.nick))
// 	}
// }

