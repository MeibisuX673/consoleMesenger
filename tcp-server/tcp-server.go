package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func processMessage(conns map[int]net.Conn, id int) {

	conn := conns[id]
	defer conn.Close()

	for {
		message := make([]byte, 1024)

		if _, err := conn.Read(message); err != nil {
			delete(conns, id)
			break
		}

		id, err := getIdConn(string(message))
		if err != nil {
			conn.Write([]byte("error format message example: id message"))
			continue
		}

		outMessage, err := getMessageText(string(message))
		if err != nil {
			conn.Write([]byte(err.Error()))
			continue
		}

		resultMessage := fmt.Sprintf("\nid: %d\nmessage: %s\n", id, outMessage)

		recipient, ok := conns[id]
		if !ok {
			conn.Write([]byte("recipient not found"))
			continue
		}

		recipient.Write([]byte(resultMessage))
	}
}

func getMessageText(message string) (string, error) {
	arr := strings.Split(message, " ")

	textMessage := strings.Join(arr[1:], " ")

	if len(textMessage) == 0 {
		return "", errors.New("text message is empty")
	}

	return textMessage, nil
}

func getIdConn(message string) (int, error) {
	arr := strings.Split(message, " ")

	id, err := strconv.Atoi(arr[0])
	if err != nil {
		return 0, err
	}

	return id, nil
}

func main() {

	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Println("Listening on :8081")

	conns := make(map[int]net.Conn)
	id := 0

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		conns[id] = conn

		log.Println("New connection: ", id, "\naddress: "+conn.RemoteAddr().String())

		conn.Write([]byte(fmt.Sprintf("your id: %d\n", id)))

		go processMessage(conns, id)

		id++
	}

}
