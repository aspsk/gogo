package main

import (
	"io"
	"fmt"
	"bufio"
	"log"
	"net"
	"os"
	"time"
)

func client(address string, clocks *string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		*clocks = string(line)
	}
}

func main() {

	n := len(os.Args[1:])
	if n == 0 {
		log.Fatal(":(")
	}
	clocks := make([]string, n)

	for i, arg := range os.Args[1:] {
		go client(arg, &clocks[i])
	}

	for {
		fmt.Printf("\033c")
		for i := 0; i < n; i++ {
			fmt.Printf("'%s' ", clocks[i])
		}

		time.Sleep(1 * time.Second)
	}
}
