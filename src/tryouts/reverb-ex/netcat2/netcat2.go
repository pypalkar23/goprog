package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// !+
func main() {
	conn, err := net.Dial("tcp", "localhost:9500")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(os.Stdin.Name())
	input := bufio.NewScanner(os.Stdin)
	defer conn.Close()
	go mustCopy(os.Stdout, conn)

	for input.Scan() {
		fmt.Fprintln(conn, input.Text())
	}

}

//!-
