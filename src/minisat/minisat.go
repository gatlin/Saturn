package minisat

/*
#include "solver.h"
#include "vec.h"
#include "minisat.h"
*/
import "C"

func Solve (text string) bool {
    res := C.solve(C.CString(text))
    var s bool
    if res == 1 {
        s = true
    } else {
        s = false
    }
    return s
}

