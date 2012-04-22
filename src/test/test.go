package main

import (
    "fmt"
    "minisat"
    "net"
    "bufio"
)

func test1() {
    fmt.Printf("TEST 1\n***\n")
    problem := "1 0\n-1 2 0\n2 0"
    solvable, _ := minisat.SolveDIMACS(problem)
    fmt.Printf("Solvable: %t\n",solvable)
}

func test2() {
    fmt.Printf("TEST 2\n***\n")
    conn,dialerr := net.Dial("tcp","localhost:1888")
    if dialerr != nil {
        fmt.Printf("ERROR dialing Saturn\n")
    }
    conn.Write([]byte("1 0\n-1 2 0\n2 0"))

    response := make([]byte,2048)
    _, reserr := bufio.NewReader(conn).Read(response)
    if reserr != nil {
        fmt.Printf("ERROR receiving response")
    }
    fmt.Printf("Response:\n%s\n",string(response))
}

func main() {
    test1()
    fmt.Printf("\n")
    test2()
}
