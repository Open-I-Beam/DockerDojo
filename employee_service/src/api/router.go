package api

import (
	"employee"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
)

var (
	employeeURL = "/employee"
)

type handler struct {
	empSvc EmployeeService
}

func (h *handler) ListEmployees(r *http.Request) responseEntity {
	emps, err := h.empSvc.ListEmployees()
	if err != nil {
		return responseEntity{http.StatusInternalServerError, err}
	} else {
		return responseEntity{http.StatusOK, emps}
	}

}

func (h *handler) AddEmployee(r *http.Request) responseEntity {
	var e employee.Employee
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		return responseEntity{http.StatusInternalServerError, err}
	}

	err := h.empSvc.AddEmployee(e)
	if err != nil {
		return responseEntity{http.StatusInternalServerError, err}
	} else {
		return responseEntity{http.StatusOK, nil}
	}
}

func (h *handler) DeleteEmployee(r *http.Request) responseEntity {
	//	vars := mux.Vars(r)

	values := r.URL.Query()
	fmt.Println("SRINI vars ", values)
	err := h.empSvc.DeleteEmployee(values.Get("Name"))
	if err != nil {
		return responseEntity{http.StatusInternalServerError, err}
	} else {
		return responseEntity{http.StatusOK, nil}
	}
}

type EmployeeService interface {
	ListEmployees() ([]employee.Employee, error)
	AddEmployee(employee.Employee) error
	DeleteEmployee(name string) error
}

type Options struct {
	Host     string
	Port     int
	Username string
	Password string
	Debug    bool
	LogFile  string
	Trace    bool
	PidFile  string
}

type router struct {
	Opts Options
	mux  *mux.Router // TODO: Replace with own simpler regexp-based mux???
}

func StartAPIServer(o Options, eSvc EmployeeService) {
	h := &handler{eSvc}
	router := newRouter(o, h)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error, 1)
	go func() {
		addr := fmt.Sprintf("%v:%v", router.Opts.Host, router.Opts.Port)
		log.Printf("Broker started: Listening at [%v]", addr)
		errCh <- http.ListenAndServe(addr, router)
	}()

	select {
	case err := <-errCh:
		log.Printf("Broker shutdown with error: %v", err)
	case sig := <-sigCh:
		var _ = sig
		log.Print("Broker shutdown gracefully")
	}

}

func newRouter(o Options, h *handler) *router {
	mux := mux.NewRouter()
	mux.Handle(employeeURL, responseHandler(h.ListEmployees)).Methods("GET")
	mux.Handle(employeeURL, responseHandler(h.AddEmployee)).Methods("PUT")
	mux.Handle(employeeURL, responseHandler(h.DeleteEmployee)).Methods("DELETE")

	return &router{o, mux}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

type responseEntity struct {
	status int
	value  interface{}
}

type responseHandler func(*http.Request) responseEntity

// Marshall the response entity as JSON and return the proper HTTP status code.
func (fn responseHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	re := fn(req)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(re.status)
	if err := json.NewEncoder(w).Encode(re.value); err != nil {
		log.Printf("Error occured while marshalling response entity: %v", err)
	}
}
