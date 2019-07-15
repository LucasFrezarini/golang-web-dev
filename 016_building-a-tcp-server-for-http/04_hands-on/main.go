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

		mux(conn, url)
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

	url, err := getURL(headerFirstLine)

	if err != nil {
		log.Println(err)
	}

	return url
}

func getScannerFirstLine(scanner *bufio.Scanner) (string, error) {
	if scanner.Scan() {
		return scanner.Text(), nil
	}

	return "", errors.New("getScannerFirstLine: No lines found")
}

func getURL(headerFirstLine string) (string, error) {
	fields := strings.Fields(headerFirstLine)

	if len(fields) > 1 {
		return strings.Fields(headerFirstLine)[1], nil
	}

	return "", errors.New(fmt.Sprintf("getURL: invalid headerFirstLine format: %s", headerFirstLine))
}

func responseMessage(conn net.Conn, message string) {
	body := message

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func mux(conn net.Conn, url string) {
	switch url {
	case "/hello":
		responseMessage(conn, "Hello World!")
		break
	case "/secret":
		responseMessage(conn, "How did you find me?")
		break
	default:
		responseMessage(conn, "Welcome. Enter on /hello to receive a hello world!")
	}
}
