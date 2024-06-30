package tcp

import (
	"bufio"
	"log"
	"net"
	"strconv"
)

var number_of_connections int = 0

func check_connection_limit() bool {
	log.Printf("number of conns: %d\n", number_of_connections)
	return number_of_connections < 2
}

func decrement_connections() {
	number_of_connections--
}

func tcp_handler(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		// Read the first byte to determine message type
		messageType, err := reader.ReadByte()
		if err != nil {
			log.Println("received error while reading, regarding client as disconnected. ", err)
			decrement_connections()
			return
		}

		switch messageType {
		case 1:
			log.Println("msg type 1 found!")
		case 2:
			log.Println("msg type 2 found!")
		default:
			log.Println("unkown message!")
			conn.Close()
			decrement_connections()
			return

		}
		conn.Write([]byte("ok"))
	}

}

func Start_tcp_server(port int) {
	log.Println("network tiger starting")

	l, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		number_of_connections += 1

		if check_connection_limit() {
			log.Println("connection accepted")
		} else {
			log.Println("connection refused, above allowed limit")
			decrement_connections()
			conn.Close()
			continue
		}

		if err != nil {
			log.Fatal(err)
		}

		go tcp_handler(conn)
	}
}
