package main

import (
	"api"
	"employee"
	"fmt"
	"jsonbuilder"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("test.txt")
	defer f.Close()
	emprecs := new(employee.EmpRecs)
	err = jsonbuilder.ReadData(f, &emprecs.Emps)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(emprecs)

	emprecs.Emps = append(emprecs.Emps, employee.Employee{"ghi", 12000, "mgr1"})
	emprecs.Sortby = func(e1, e2 *employee.Employee) bool {
		return e1.Salary <= e2.Salary
	}

	//TODO add the implementor of the EmployeeService
	api.StartAPIServer(api.Options{Host: "localhost", Port: 5678}, interface{})

	sort.Sort(emprecs)
	f2, err := os.Create("test2.txt")
	defer f2.Close()
	jsonbuilder.WriteData(f2, emprecs.Emps)
	fmt.Println(emprecs)
}
