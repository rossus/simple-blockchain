package tcp

import (
	"log"
	"net"
	"os"
)

func Run() {
	server, err := net.Listen("tcp", ":"+os.Getenv("ADDR2"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on ", os.Getenv("ADDR2"))
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}