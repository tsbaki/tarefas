package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "bytes"
    "fmt"
    "json"
)

func TestHandlerHealthCheck(t *testing.T) {
    req, err := http.NewRequest("GET", "/health-check", nil)
    if err != nil {
        t.Fatal()
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HandlerHealthCheck)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("HealthCheck returned wrong status code: %v", 
        status)
    }

    expected := `{"dbExists": true}`
    if rr.Body.String() != expected {
        t.Errorf("HealthCheck returned unexpected body: %v; Expected %v.", 
        rr.Body.String(), expected)
    }
}

func TestHandlerAddTodo(t *testing.T) {
    var b bytes.Buffer
    var todo Todo
    todo.Project = "ProjectA"
    todo.Description = "This is a test todo"

    err := json.NewEncoder(b).Encode(&t)
    req, rerErr := http.NewRequest(http.MethodPost, "/addTodo", nil)

    if (rerErr != nil) {
        t.Fatal()
        panic(reqErr)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HandlerAddTodo)
    handler.ServeHTTP(rr, req)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Couldn't add todo, status code was: %v", 
        status)
    }
}

func TestHandlerTodos(t *testing.T) {
    req, err := http.NewRequest("GET", "/todos", nil)

    if err != nil {
        t.Fatal()
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HandlerHealthCheck)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("Bad status code returned %v", status)
    }

    if rb := rr.Body.String(); len(rb) == 0 {
        t.Errorf("Test failed: Nothing returned")
        t.Fail()
    } else  {
        fmt.Println(rb)
    }
}
