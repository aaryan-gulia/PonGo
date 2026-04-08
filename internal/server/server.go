package server

import (
	"PonGo/internal/pong"
	"bytes"
	"encoding/gob"
	"log"
	"net"
)

type Game struct {
	state   pong.GameState
	addr    []*net.UDPAddr
	conn    *net.UDPConn
	chState chan pong.GameState
}

func Run() {
	conn := setup()
	defer conn.Close()

	buffer := make([]byte, 1024)
	var addr []*net.UDPAddr
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
			addr = append(addr, clientAddr2)
			break
		}
	}

	runGame(addr, conn)
}

func runGame(addr []*net.UDPAddr, conn *net.UDPConn) {
	game := Game{addr: addr, conn: conn, chState: make(chan pong.GameState)}

	game.state.Reset()

	if len(game.addr) == 2 {
		log.Println("starting game")
	} else {
		log.Println("not enough players connected")
	}

	for {
		game.state.PollState()

		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		if err := enc.Encode(game.state); err != nil {
			log.Fatalln("encoding error : ", err)
		}

		_, err := conn.WriteToUDP(buf.Bytes(), game.addr[0])
		if err != nil {
			log.Println("writing error : ", err)
		}
		_, err = conn.WriteToUDP(buf.Bytes(), game.addr[1])
		if err != nil {
			log.Println("writing error : ", err)
		}

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
