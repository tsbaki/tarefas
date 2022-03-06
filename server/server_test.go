package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHandlerHealthCheck(t *testing.T) {
    req, err := http.NewRequest("GET", "/health-check", nil)
    if err != nil {
        t.Fatal()
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlerHealthCheck)
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

func TestHandlerTodos(t *testing.T) {
    t.Log("Test not implemented yet.")
}
