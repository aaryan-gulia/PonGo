package server

import (
	"PonGo/internal/client"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

func Run() {
	conn := setup()
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("read error: ", err)
		}

		var e client.GameEvent
		dec := gob.NewDecoder(bytes.NewReader(buffer[:n]))
		if err := dec.Decode(&e); err != nil {
			log.Panicln("decoding error : ", err)
		}
		fmt.Println("event: ", e)
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
