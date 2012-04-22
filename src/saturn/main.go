package main

import (
	"fmt"
	"net"
	"minisat"
	"bytes"
)

func main() {
	/* listen on a port for incoming problems */
	ln, err := net.Listen("tcp", ":1888")
	if err != nil {
		fmt.Printf("ERROR on listen\n")
	}

	defer ln.Close()

	/* each request gets a goroutine */
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("ERROR on accept")
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	buffer := make([]byte, 2048)

	_, err := c.Read(buffer)
	if err != nil {
		fmt.Printf("ERROR on read\n")
	}

	solvable, solution := minisat.SolveDIMACS(string(buffer))
	if solvable {
		c.Write(writeSolution(solution))
	} else {
		c.Write([]byte("0\n"))
	}
	c.Close()
}

func writeSolution(b []bool) []byte {
	buffer := bytes.NewBufferString("")
	fmt.Fprint(buffer, "1\n")
	for variable, val := range b {
		var t int
		if val {
			t = 1
		} else {
			t = 0
		}
		fmt.Fprintf(buffer, "%d %d\n", variable+1, t)
	}
	return buffer.Bytes()
}
