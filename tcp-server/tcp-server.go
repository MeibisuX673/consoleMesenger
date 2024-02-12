package main

import (
	"errors"
	"fmt"
	"github.com/MeibisuX673/consoleMesenger/tcp-server/beats/fileServerBeat"
	"log"
	"net"
	"strconv"
	"strings"
)

var Version string
var User string
var Date string

func processMessage(conns map[int]net.Conn, id int) {

	conn := conns[id]
	defer conn.Close()

	for {
		message := make([]byte, 1024)

		if _, err := conn.Read(message); err != nil {
			delete(conns, id)
			break
		}

		idTo, err := getIdConn(string(message))
		if err != nil {
			conn.Write([]byte("error format message example: id message"))
			continue
		}

		recipient, ok := conns[idTo]
		if !ok {
			conn.Write([]byte("recipient not found"))
			continue
		}

		outMessage, err := getMessageText(string(message))
		if err != nil {
			conn.Write([]byte(err.Error()))
			continue
		}

		resultMessage := fmt.Sprintf("\nid: %d\nmessage: %s\n", id, outMessage)

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

	fileServerBeat.CreateLog("Run Server Listening on :8081")

	log.Println("Listening on :8081")

	conns := make(map[int]net.Conn)
	id := 0

	printInfo()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		conns[id] = conn

		fileServerBeat.CreateLog(fmt.Sprintf("New connection: %d address: %s", id, conn.RemoteAddr().String()))

		_, err = conn.Write([]byte(fmt.Sprintf("your id: %d\n", id)))
		if err != nil {
			fileServerBeat.CreateLog(fmt.Sprint("error connection: ", id, "\naddress: "+conn.RemoteAddr().String()))

		}

		go processMessage(conns, id)

		id++
	}
}

func printInfo() {
	fmt.Println("Version: " + Version)
	fmt.Println("User: " + User)
	fmt.Println("Date: " + Date)
}
