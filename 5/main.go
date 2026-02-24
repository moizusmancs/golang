package main

import (
	"bufio"
	"fmt"
	"net"
)

func main(){


	listener, error := net.Listen("tcp", ":8080")

	if error != nil {
		panic(error)
	}

	defer listener.Close()

	fmt.Println("Server listeneing on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		conn.Write([]byte("Welcome to our server!"))
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()

	fmt.Println("New client:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			return
		}
		fmt.Print("Received:", message)

		conn.Write([]byte("Echo: " + message))
	}

}