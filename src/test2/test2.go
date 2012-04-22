package main

import (
    "fmt"
    "net"
    "bufio"
)

func main() {
    conn,dialerr := net.Dial("tcp","localhost:1888")
    if dialerr != nil {
        fmt.Printf("well shit\n")
    }
    conn.Write([]byte("1 0\n-1 2 0\n2 0"))

    response := make([]byte,2048)
    _, reserr := bufio.NewReader(conn).Read(response)
    if reserr != nil {
        fmt.Printf("ERROR receiving response")
    }
    fmt.Printf("Response: %s\n",string(response))
}
