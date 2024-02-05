package main

import (
	"bufio"
	"fmt"
	"github.com/MeibisuX673/consoleMesenger/tcp-client/config"
	"github.com/MeibisuX673/consoleMesenger/tcp-client/config/connection"
	"log"
	"net"
	"os"
)

var Version string
var User string
var Date string

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
			printMessage("disconnect")
			disconnect <- true
			return
		}

		printMessage("\n" + string(message))
	}
}

func printMessage(message string) {

	fmt.Println(config.Flags.UseColor + message)
}

func main() {

	connection.GetConnection()
	conn, err := net.Dial("tcp", connection.GetConnection())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	disconnect := make(chan bool)

	go readConsole(conn)
	go readSocket(conn, disconnect)

	printInfo()

	<-disconnect
}

func printInfo() {
	printMessage("Version: " + Version)
	printMessage("User: " + User)
	printMessage("Date: " + Date)
}
