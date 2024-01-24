package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func readConsole(
	conn net.Conn,
) {
	var message string

	reader := bufio.NewReader(os.Stdin)

	for {
		message = ""

		message, _ = reader.ReadString('\n')

		if _, err := conn.Write([]byte(message)); err != nil {
			log.Fatal(err)
		}
	}
}

func readSocket(
	conn net.Conn,
	disconnect chan bool,
) {
	for {
		message := make([]byte, 1024)
		if _, err := conn.Read(message); err != nil {
			fmt.Println("disconect")
			disconnect <- true
			return
		}

		fmt.Print("\n" + string(message))
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	disconnect := make(chan bool)

	go readConsole(conn)
	go readSocket(conn, disconnect)

	<-disconnect

}
