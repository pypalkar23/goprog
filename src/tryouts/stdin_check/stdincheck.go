package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Fprintln(os.Stdout, "helloworld")
	//reader := bufio.NewReader(os.Stdin)
	//input, _ := reader.ReadString('\n')
	writer := bufio.NewWriter(os.Stdout)
	if _, err := io.Copy(writer, os.Stdin); err != nil {
		fmt.Println("Error while writing")
	}
}
