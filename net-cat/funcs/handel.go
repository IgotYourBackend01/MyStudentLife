package funcs

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func HandelConn(conn net.Conn) {
	timeClient := time.Now().Format("2006-01-02 15:04:05")
	Welcom(conn)
	nickname := GetName(conn)
	tempClient := Client{
		name: nickname,
		addr: conn.RemoteAddr().String(),
		conn: conn,
	}
	mutex.Lock()
	clients[nickname] = tempClient
	if len(clients) > 10 {
		mutex.Lock()
		tempClient.conn.Write([]byte("Chat is full!"))
		delete(clients, nickname)
		tempClient.conn.Close()
		mutex.Unlock()
		return
	}
	mutex.Unlock()
	mutex.Lock()
	join <- NewMessege("has joined our chat...", conn, tempClient, timeClient)
	mutex.Unlock()
	fmt.Fprintf(conn, "[%s][%s]:", timeClient, nickname)
	input := bufio.NewScanner(conn)

	for input.Scan() {
		time := time.Now().Format("2006-01-02 15:04:05")
		if input.Text() == "" {
			fmt.Fprintf(conn, "you can't send empty messages\n")
			fmt.Fprintf(conn, "[%s][%s]:", time, nickname)
			continue
		}
		if !IsCorrect(input.Text(), conn, time, nickname) {
			fmt.Fprintf(conn, "Incorrext inptu\n")
			fmt.Fprintf(conn, "[%s][%s]:", time, nickname)
			continue
		}
		text := fmt.Sprintf("[%s][%s]:%s\n", time, nickname, input.Text())
		mutex.Lock()
		historytext = append(historytext, text)
		messages <- NewMessege(input.Text(), conn, tempClient, time)
		mutex.Unlock()
	}
	mutex.Lock()
	delete(clients, nickname)
	leaving <- NewMessege("has left our chat...", conn, tempClient, timeClient)
	conn.Close()
	mutex.Unlock()
}
