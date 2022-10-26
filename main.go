package main

import (
	"fmt"

	"net"
)

func main() {
	message := "Hello from my server(it was sent from server)"
	listener, err := net.Listen("tcp", ":4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening!")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte(message))
		conn.Close()
	}

}
