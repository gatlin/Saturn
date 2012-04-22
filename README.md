Saturn
===

A distributed SAT solver written in Go.
(c) 2012, Gatlin Johnson <rokenrol@gmail.com>

0. What is Saturn?
---

Saturn is a distributed SAT solver based on [MiniSat][1]. Its aims are to be

* distributed
* scalable
* fault tolerant

1. How does it work?
---

A single Saturn instance is aware of other instances on the same network using
multicast. Configuration is automatic. It listens on the network for problems in
the [DIMACS][2] format.

Then, the problem is split up by making intelligent guess assignments for
certain literals; each of these is sent to an arbitrary instance to be worked
on independently. In this way, if there are *N* nodes then a problem can, in
principle, be split into *N* sub problems and searched in parallel.

2. What are the advantages of Saturn over other solvers?
---

Saturn is based on MiniSAT, though I hope to add more solving backends in the
future to take a portfolio approach. Saturn's aim is to make it painless to
create large clusters and split up the search space of larger problems.

3. How do I build it?
---

1. Make sure Go 1 is installed.
2. `cd` to where this repository is checked out.
3. `export GOPATH=$PWD/`
4. `go install saturn`

Saturn will be installed and built in the `bin/` directory. You can also `go
install test` to install the test program.

4. Progress?
---

* 22 April 2012: Saturn now has an output format and returns that
* 21 April 2012: Networked SAT oracle! (view history)
* 15 April 2012: Bindings are created; not sure they're optimal.
* 12 April 2012: This readme was created.

[1]: http://minisat.se/
[2]: http://www.satcompetition.org/2009/format-benchmarks2009.html
