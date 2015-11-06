package employee

import (
	"errors"
	"fmt"
)

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

func (e *EmpRecs) AddEmployee(emp Employee) error {
	e.Emps = append(e.Emps, emp)
	return nil
}

func (e *EmpRecs) DeleteEmployee(name string) error {
	for i, emp := range e.Emps {
		if emp.Name == name {
			newSlice := make([]Employee, i)
			copy(newSlice, e.Emps[:i])
			if i < len(e.Emps) {
				newSlice = append(newSlice, e.Emps[i+1:]...)
			}
			e.Emps = newSlice
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Not Found %s", name))
}

func (e *EmpRecs) ListEmployees() ([]Employee, error) {
	return e.Emps, nil
}
