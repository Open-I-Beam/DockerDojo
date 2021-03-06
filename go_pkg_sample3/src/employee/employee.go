package employee

import ()

type Employee struct {
	Name    string
	Salary  float32
	Manager string
}

type EmpRecs struct {
	Emps   []Employee
	Sortby func(p1, p2 *Employee) bool
}

func (e *EmpRecs) Len() int {
	return len(e.Emps)
}

func (e *EmpRecs) Less(i, j int) bool {
	return e.Sortby(&e.Emps[i], &e.Emps[j])
}

func (e *EmpRecs) Swap(i, j int) {
	e.Emps[i], e.Emps[j] = e.Emps[j], e.Emps[i]
}
