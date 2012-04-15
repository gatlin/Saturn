package minisat

/*
#include "solver.h"
#include "vec.h"
#include "minisat.h"
*/
import "C"

func SolveDIMACS(text string) (solvable bool, soln []bool) {
	slv := C.make_solver()
	res := C.solve(slv, C.CString(text))
	if res == 1 {
		solvable = true
		soln = solution(slv)
	} else {
		solvable = false
		soln = nil
	}
	C.solver_delete(slv)
	return solvable, soln
}

func solution(slv *C.solver) []bool {
	size := C.assignment_size(slv)
	var assignment []bool = make([]bool, size)
	for i := 0; C.int(i) < size; i++ {
		if C.get_assignment(slv, C.int(i)) == 1 {
			assignment[i] = true
		} else {
			assignment[i] = false
		}
	}
	return assignment
}
