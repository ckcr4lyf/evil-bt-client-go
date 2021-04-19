package sender

import (
	"log"
	"net"

	"github.com/ckcr4lyf/evil-bt-client-go/internal/message"
)

func SendData() {
	conn, err := net.Dial("tcp", "127.0.0.1:1337")

	if err != nil {
		log.Print("Failed to make TCP conenction", err)
	}

	//TODO: Seed with hash, number of pieces

	conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))

	recvBuffer := make([]byte, 1024)

	_, err = conn.Read(recvBuffer)

	if err != nil {
		log.Print("Failed to read from TCP socket")
	}

	log.Print(string(recvBuffer))
	log.Print(message.MakeHandshakeMessage())
}
