package main

import (
	"fmt"
	"net-cat/funcs"
	"os"
)

func main() {
	port := ":8989"
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	} else if len(os.Args) == 2 {
		port = ":" + os.Args[1]
	}
	funcs.CreateServerPort(port)
}
