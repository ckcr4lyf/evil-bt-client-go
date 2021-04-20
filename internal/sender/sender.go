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
	recvBuffer := make([]byte, 68)
	_, err = conn.Read(recvBuffer)

	if err != nil {
		log.Print("Failed to read from TCP socket")
	}

	// log.Printf("% x", recvBuffer[4:binary.BigEndian.Uint32(recvBuffer[0:4])])
	log.Printf("The handshake response is % x", recvBuffer)
	conn.Write(message.MakeInterestedMessage()) // Send an interested request
	recvBuffer = make([]byte, 1024)
	_, err = conn.Read(recvBuffer)

	if err != nil {
		log.Print("Failed to read from TCP socket")
	}

	// log.Printf("% x", recvBuffer)
	log.Printf("% x", recvBuffer[0:4])
	ps := binary.BigEndian.Uint32(recvBuffer[0:4])
	log.Print("The payload size is ", ps)
	log.Printf("The payload is % x", recvBuffer[4:4+ps])
	log.Print(recvBuffer)
	// os.Exit(0)

	conn.Write(message.MakeRequest(0, 0)) // Send a piece request
	recvBuffer = make([]byte, 4)
	_, err = conn.Read(recvBuffer)

	if err != nil {
		log.Print("Failed to read from TCP socket")
		log.Fatal(err)
	}

	log.Printf("% x", recvBuffer)
	payloadSize := binary.BigEndian.Uint32(recvBuffer)
	recvBuffer = make([]byte, payloadSize+1)
	_, err = conn.Read(recvBuffer)

	if err != nil {
		log.Print("Failed to read from TCP socket")
		log.Fatal(err)
	}

	log.Printf("% x", recvBuffer)

	// log.Printf("% x", message.MakeHandshakeMessage("28dc4be125a69329b7dd1d90880cd2e3c514b384"))
}
