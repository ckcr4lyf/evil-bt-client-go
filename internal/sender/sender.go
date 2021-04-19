package sender

import (
	"encoding/binary"
	"log"
	"net"
	"os"

	"github.com/ckcr4lyf/evil-bt-client-go/internal/message"
)

func SendData() {

	args := os.Args

	if len(args) < 3 {
		log.Fatal("Insufficient number of arguments")
	}

	if len(args[2]) != 40 {
		log.Fatal("Infohash must be 40 characters long")
	}

	conn, err := net.Dial("tcp", args[1])

	if err != nil {
		log.Print("Failed to make TCP conenction", err)
	}

	//TODO: Seed with hash, number of pieces

	// Send the handshake
	handshakeMessage := message.MakeHandshakeMessage(args[2])
	conn.Write(handshakeMessage)

	recvBuffer := make([]byte, 1024)

	_, err = conn.Read(recvBuffer)

	if err != nil {
		log.Print("Failed to read from TCP socket")
	}

	log.Print(string(recvBuffer))
	log.Print(len(recvBuffer))

	// Send an interested request
	conn.Write(message.MakeInterestedMessage())

	recvBuffer = make([]byte, 1024)

	_, err = conn.Read(recvBuffer)

	if err != nil {
		log.Print("Failed to read from TCP socket")
	}

	log.Print(string(recvBuffer))
	log.Print(len(recvBuffer))

	// Send a piece request
	conn.Write(message.MakeRequest(0, 0))

	recvBuffer = make([]byte, 4)
	_, err = conn.Read(recvBuffer)

	if err != nil {
		log.Print("Failed to read from TCP socket")
		log.Fatal(err)
	}

	log.Printf("% x", recvBuffer)
	log.Print(len(recvBuffer))

	payloadSize := binary.BigEndian.Uint32(recvBuffer)
	recvBuffer = make([]byte, payloadSize+1)

	_, err = conn.Read(recvBuffer)

	if err != nil {
		log.Print("Failed to read from TCP socket")
		log.Fatal(err)
	}

	log.Printf("% x", recvBuffer)
	log.Print(len(recvBuffer))

	// log.Printf("% x", message.MakeHandshakeMessage("28dc4be125a69329b7dd1d90880cd2e3c514b384"))
}
