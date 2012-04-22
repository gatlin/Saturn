package main

import (
    "fmt"
    "net"
    "minisat"
)

func main() {
    /* listen on a port for incoming problems */
    ln, err := net.Listen("tcp",":1888")
    if err != nil {
        fmt.Printf("ERROR on listen\n")
    }

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
    buffer := make([]byte,2048)

    _, err := c.Read(buffer)
    if err != nil {
        fmt.Printf("ERROR on read\n")
    }

    solvable, _ := minisat.SolveDIMACS(string(buffer))
    if solvable {
        c.Write([]byte("Solvable"))
    } else {
        c.Write([]byte("Not solvable"))
    }
    c.Close()
}

