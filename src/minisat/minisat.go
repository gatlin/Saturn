package minisat

/*
#include "solver.h"
#include "vec.h"
#include "minisat.h"
*/
import "C"

func SolveDIMACS (text string) bool {
    res := C.solve(C.CString(text))
    var status bool

    if res == 1 {
        status = true
    } else {
        status = false
    }

    return status
}

func Solution() []bool {
    size := C.assignment_size()
    var assignment []bool = make([]bool,size)
    for i := 0; C.int(i) < size; i++ {
        if C.get_assignment(C.int(i)) == 1 {
            assignment[i] = true
        } else {
            assignment[i] = false
        }
    }
    return assignment
}

func Finish() {
    C.solver_delete(C.get_solver())
}

