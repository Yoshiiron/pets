package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

// struct for Clients in chat app.
type Client struct {
	conn net.Conn
	name string
}

// vars init
// mutex for access to clients
var (
	mu      sync.Mutex
	clients = make(map[string]Client)
)

// Func for handle clients.
func clinetHandler(conn net.Conn) {
	defer conn.Close() //close connection when exit the func

	// Get name of the client
	conn.Write([]byte("Enter your name: "))
	nameScanner := bufio.NewScanner(conn)
	nameScanner.Scan()
	name := strings.TrimSpace(nameScanner.Text())

	client := Client{conn: conn, name: name}

	// add client to the map
	mu.Lock()
	clients[name] = client
	mu.Unlock()

	// entering message
	conn.Write([]byte(fmt.Sprintln("Welcome! You can use /msg [username] [message] to write.")))

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		if strings.HasPrefix(msg, "/msg") {
			// split the command
			parts := strings.SplitN(msg, " ", 3)
			// check if the entered message is correct
			if len(parts) < 3 {
				conn.Write([]byte(fmt.Sprintln("Invalid format, please use '/msg [username] [message]'")))
				continue
			}
			// parse name and message
			recipientName := parts[1]
			msg := parts[2]

			// check if nickname that entered in msg is exist in clients map
			mu.Lock()
			recipient, exists := clients[recipientName]
			mu.Unlock()

			if exists {
				recipient.conn.Write([]byte(fmt.Sprintf("%s: %s\n", name, msg)))
			} else {
				conn.Write([]byte(fmt.Sprintln("User not found")))
			}
			// if user entered /quit - exit from app.
		} else if strings.HasPrefix(msg, "/quit") {
			conn.Write([]byte(fmt.Sprintln("Goodbye.")))
			break
		} else {
			conn.Write([]byte(fmt.Sprintln("Use /msg <username> <message> to send private messages")))
		}
	}

	// delete client from map
	mu.Lock()
	delete(clients, name)
	mu.Unlock()

}

func main() {
	// create a server
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	// defer for closing a server
	defer ln.Close()

	for {
		//accept the conn
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// handle the client
		go clinetHandler(conn)
	}
}
