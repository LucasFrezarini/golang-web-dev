package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	server := createServer()
	defer server.Close()

	for {
		conn := accept(server)
		defer conn.Close()

		url := getURLFromRequest(conn)
		fmt.Println(url)
	}
}

func createServer() net.Listener {
	server, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	return server
}

func accept(server net.Listener) net.Conn {
	conn, err := server.Accept()

	if err != nil {
		log.Println(err)
	}

	return conn
}

func getURLFromRequest(conn net.Conn) string {
	scanner := bufio.NewScanner(conn)
	headerFirstLine, err := getScannerFirstLine(scanner)

	if err != nil {
		log.Println(err)
	}

	return getURL(headerFirstLine)
}

func getScannerFirstLine(scanner *bufio.Scanner) (string, error) {
	if scanner.Scan() {
		return scanner.Text(), nil
	}

	return "", errors.New("getScannerFirstLine: No lines found")
}

func getURL(headerFirstLine string) string {
	return strings.Fields(headerFirstLine)[1]
}
