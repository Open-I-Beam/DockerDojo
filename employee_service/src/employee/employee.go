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
	//TODO write code to add new employee
	return nil
}

func (e *EmpRecs) DeleteEmployee(name string) error {
	//TODO write code to delete the employee from the list
	return errors.New(fmt.Sprintf("Not Found %s", name))
}

func (e *EmpRecs) ListEmployees() ([]Employee, error) {
	//TODO write code to return list of employees
	return nil, nil
}
