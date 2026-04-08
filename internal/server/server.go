package server

import (
	"PonGo/internal/pong"
	"log"
	"net"
)

type Game struct {
	state   pong.GameState
	addr    []net.Addr
	conn    *net.UDPConn
	chState chan pong.GameState
}

func Run() {
	conn := setup()
	defer conn.Close()

	buffer := make([]byte, 1024)
	var addr []net.Addr
	_, clientAddr1, err := conn.ReadFromUDP(buffer)
	if err != nil {
		log.Println("read error: ", err)
	}

	addr = append(addr, clientAddr1)
	log.Println("client address:", clientAddr1)
	for {
		_, clientAddr2, err := conn.ReadFromUDP(buffer)
		log.Println("client address:", clientAddr2)
		if err != nil {
			log.Println("read error: ", err)
		}
		if clientAddr2.String() != clientAddr1.String() {
			addr = append(addr, clientAddr1)
			break
		}
	}

	runGame(addr, conn)
}

func runGame(addr []net.Addr, conn *net.UDPConn) {
	game := Game{addr: addr, conn: conn, chState: make(chan pong.GameState)}

	game.state.Reset()

	if len(game.addr) == 2 {
		log.Println("starting game")
	} else {
		log.Println("not enough players connected")
	}

}

func setup() *net.UDPConn {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8080")

	if err != nil {
		log.Fatal("could not resolve address: ", err)
	}

	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		log.Fatal("listen failed: ", err)
	}

	return conn
}
