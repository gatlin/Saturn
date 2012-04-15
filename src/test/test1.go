package main

import (
    "fmt"
    "minisat"
)

func main() {
    problem := "1 0\n-1 2 0\n2 0"
    solvable,solution := minisat.SolveDIMACS(problem)
    if solvable {
        for k,v := range solution {
            fmt.Printf("var %d = %t ",k,v)
        }
        fmt.Printf("\n")
    } else {
        fmt.Printf("No\n")
    }
}
