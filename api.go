package main

import (
    "os"
    "log"
    "net/http"
    "io"
    "io/ioutil"
    "fmt"
    "encoding/json"
)

const DB_PATH   = "todos.json"

    func HandlerInit(w http.ResponseWriter, r *http.Request) { 
        if r.Method != "GET" { return }
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "/todos - Retrieves all todos\n")
    }

    /*
    * Preforms a basic status health check and if the server
    * actually contains a database ready to store our todos.
    */
    func HandlerHealthCheck(w http.ResponseWriter, r *http.Request) {
        // To be used for making sure a database actually
        // exists.
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Contect-Type", "application/json")
        //TODO: Allow for the user to set an env variable
        //to override the default path.
        //TODO: Refactor this.
        if home, err := os.UserHomeDir(); err == nil  {
            if _, err := os.Open(home + "/.tarefas"); err == nil { 
                io.WriteString(w, `{"dbExists": true}`)
            } else {
                io.WriteString(w, `{"dbExists": false}`)
            }
        } else {
            io.WriteString(w, `{"dbExists": false}`)
        }
    }

    /*
    * Gets our list of TODOS
    * Sends a response back as a list of JSON objects
    */
    func HandlerTodos(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" { return }
    }

    func HandlerAddTodo(w http.ResponseWriter, r *http.Request) { 
        if r.Method != "POST" { return }
        body, err := ioutil.ReadAll(r.Body)
        if err != nil { 
            log.Println("[ERROR]: Couldn't req body: %s", err)
        } else {
        var t Todo
        if err = json.Unmarshal(body, &t); err != nil { 
            log.Println("[ERROR]: Couldn't unmarshal object: %s", 
            err) 
        } else {
            AddTodo(&t)
        }
    }

}
