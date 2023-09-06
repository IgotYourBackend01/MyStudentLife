package funcs

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func NewMessege(msg string, conn net.Conn, cl Client, time string) Message {
	return Message{
		text:    msg,
		address: cl.addr,
		name:    cl.name,
		time:    time,
		history: historytext,
	}
}

func Welcom(conn net.Conn) {
	file, err := os.ReadFile("logo.txt")
	if err != nil {
		fmt.Printf("couldn't read this file")
	}
	imagetxt := string(file)
	conn.Write([]byte("Welcom to TCP-Chat!\n" + imagetxt + "\n"))
}

func GetName(conn net.Conn) string {
	conn.Write([]byte("[Enter your name]:"))
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}
	temp := strings.TrimSpace(string(data))

	if temp == "" || len(temp) == 0 || len(temp) > 20 {
		fmt.Fprintln(conn, "Incorrect input")
		return GetName(conn)
	}

	for _, i := range temp {
		if i < 32 && i > 127 {
			fmt.Fprintf(conn, "Incorrect input\n")
			return GetName(conn)
		}
	}
	for client := range clients {
		if client == temp {
			fmt.Fprintf(conn, "User already exist\n")
			return GetName(conn)
		}
	}
	return temp
}

// функция правильности ввода каждого подключенного к серверу пользователя
func IsCorrect(s string, conn net.Conn, time string, username string) bool {
	for _, w := range s {
		if w < 32 || w > 127 {
			return false
		}
	}
	return true
}
