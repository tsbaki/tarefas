package main

import (
    "log"
    "io/ioutil"
    "net/http"
    "fmt"
    "encoding/json"
    )

var (
    //TODO(andre): turn this into a environment
    //variable.
    port    = "8080"
    todos   = []todo_t{ }
)

type todo_t struct {
    Description, Project string
}

func initialHandler(w http.ResponseWriter, r *http.Request) { 
    if r.Method != "GET" {
        return;
    }
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "/todos - Retrieves all todos\n")
}

func handlerTodos(w http.ResponseWriter, r *http.Request) {
}

func handlerAddTodo(w http.ResponseWriter, r *http.Request) { 
    if r.Method != "POST" {
        return;
    }
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err)
    }
    var t todo_t
    err = json.Unmarshal(body, &t)
    if err != nil {
        panic(err)
    }
    log.Printf("NEW TODO: '%s'; FROM '%s'", t.Description, t.Project);
    todos = append(todos, t)
}

func main() {
    log.Print("Listenning on port: " + port);
    http.HandleFunc("/", initialHandler)
    http.HandleFunc("/todos", handlerTodos)
    http.HandleFunc("/todo", handlerAddTodo)
    log.Fatal(http.ListenAndServe(":" + string(port), nil));
}
