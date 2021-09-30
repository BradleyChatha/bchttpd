package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	tree *RadixTree
)

func main() {
	tree = createRadixTreeFromDir(os.Getenv("BCHTTPD_ROOT"))
	tree.print()
	startServer()
}

func startServer() {
	log.Println("Starting server")

	server, err := net.Listen("tcp", ":"+os.Getenv("BCHTTPD_PORT"))
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Printf("Could not accept connection: %v\n", err)
			continue
		}

		go handleConnect(conn)
	}
}

func handleConnect(conn net.Conn) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	io := bufio.NewReadWriter(reader, writer)

	for {
		line, err := io.ReadString('\n')
		if err != nil {
			return
		}

		method := ""
		path := ""

		if line[0] == 'G' {
			method = "GET"
		} else {
			fmt.Println("Only GET is supported for now.")
			continue
		}

		path = line[len(method)+1 : len(line)-len(" HTTP/1.1\r\n")]
		contents := tree.find(path)

		for line != "\r\n" {
			line, err = io.ReadString('\n')
			if err != nil {
				return
			}
		}

		if len(contents) == 0 {
			_, err = io.WriteString("HTTP/1.1 404 NOT FOUND\r\nContent-Length: 9\r\n\r\nNot Found")
			if err != nil {
				return
			}
		} else {
			_, err = io.WriteString(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %v\r\n\r\n%v", len(contents), string(contents)))
			if err != nil {
				return
			}
		}
		err = io.Flush()
		if err != nil {
			return
		}
	}
}
