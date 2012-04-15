#ifndef _MINISAT_H
#define _MINISAT_H

int solve(solver *,char *);
solver *make_solver();
int assignment_size(solver *);
int get_assignment(solver *,int);

#endif
