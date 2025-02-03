package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}
	go mustCopy(os.Stdout, conn)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Fprintln(conn, input.Text())
	}j
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
