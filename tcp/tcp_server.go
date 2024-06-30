package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

var number_of_connections int = 0

func check_connection_limit() bool {
	fmt.Printf("number of conns: %d\n", number_of_connections)
	return number_of_connections < 2
}

func tcp_handler(conn net.Conn) {
	reader := bufio.NewReader(conn)

	for {
		// Read the first byte to determine message type
		messageType, err := reader.ReadByte()

		if err != nil {
			log.Fatal(err)
		}

		switch messageType {
		case 1:
			fmt.Println("msg type 1 found!")
		case 2:
			fmt.Println("msg type 2 found!")
		default:
			fmt.Println("unkown message!")
			conn.Close()
			number_of_connections--
			return

		}
		conn.Write([]byte("ok"))
	}

}

func Start_tcp_server(port int) {
	fmt.Println("network tiger starting")

	l, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		number_of_connections += 1

		if check_connection_limit() {
			fmt.Println("connection accepted")
		} else {
			fmt.Println("connection refused, above allowed limit")
			conn.Close()
		}

		if err != nil {
			log.Fatal(err)
		}

		go tcp_handler(conn)
	}
}
