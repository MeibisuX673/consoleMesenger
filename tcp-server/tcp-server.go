package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

var Version string
var User string
var Date string

var LovesPeople []string = []string{
	"Виталя",
	"Андрей",
	"Семен",
}

func CheckLoveMessage(message []byte, conn net.Conn) {
	messageString := string(message)

	isLove := strings.Contains(messageString, "love")
	if !isLove {
		return
	}

	for _, name := range LovesPeople {

		if strings.Contains(messageString, name) && isLove {
			sendLove(conn, name)
			break
		}
	}
}

func processMessage(conns map[int]net.Conn, id int) {

	conn := conns[id]
	defer conn.Close()

	for {
		message := make([]byte, 1024)

		if _, err := conn.Read(message); err != nil {
			delete(conns, id)
			break
		}

		CheckLoveMessage(message, conn)

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

func sendLove(conn net.Conn, name string) {
	love := fmt.Sprintf(""+
		"\t\t     lovelove\t\t            lovelove"+
		"\n\t\t  lovelovelovelove             lovelovelovelove"+
		"\n\t      lovelovelovelovelovelove     lovelovelovelovelovelove"+
		"\n\t    lovelovelovelovelovelovelovelovelovelovelovelovelovelove"+
		"\n\t    lovelovelovelovelovelovelovelovelovelovelovelovelovelove"+
		"\n\t    lovelovelovelovelovelovel%slovelovelovelovelovelove"+
		"\n\t      lovelovelovelovelovelovelovelovelovelovelovelovelove"+
		"\n\t        lovelovelovelovelovelovelovelovelovelovelovelove"+
		"\n\t          lovelovelovelovelovelovelovelovelovelovelove"+
		"\n\t              lovelovelovelovelovelovelovelovelove"+
		"\n\t                  lovelovelovelovelovelovelove"+
		"\n\t                      lovelovelovelovelove"+
		"\n\t                          lovelovelove"+
		"\n\t                              love\n", name)
	conn.Write([]byte(love))
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

	ln, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Println("Listening on :8083")

	conns := make(map[int]net.Conn)
	id := 0

	printInfo()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		conns[id] = conn

		conn.Write([]byte(fmt.Sprintf("your id: %d\n", id)))

		go processMessage(conns, id)

		id++
	}

}

func printInfo() {
	fmt.Println("Version: " + Version)
	fmt.Println("User: " + User)
	fmt.Println("Date: " + Date)
}
