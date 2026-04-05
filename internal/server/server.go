package server

import (
	"PonGo/internal/pong"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func Run() {
	conn := setup()
	defer conn.Close()
	var game pong.GameState

	buffer := make([]byte, 1024)
	for {
		n, ClientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("read error: ", err)
		}
		fmt.Println("message: ", string(buffer[:n]))
		event := strings.TrimSpace(string(buffer[:n]))
		if event == "w" {
			game.MovePaddle1Up()
		}
		if event == "s" {
			game.MovePaddle1Down()
		}
		log.Println("unknown event: ", event)
		conn.WriteToUDP([]byte(strconv.Itoa(int(game.Paddle1))), ClientAddr)
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
