package main

import (
    "os/exec"
    "os"
    "runtime"
    "log"
    "net/http"
    "fmt"
)

const LOG_PATH  = ".tarefas_log"
const PORT      = ":8080"

var clear map[string]func()

/*
* Creates a logging file to log any captured errors
*/
func setupFileLogging() {
    os.Remove(LOG_PATH)
    file, err := os.OpenFile(LOG_PATH, os.O_APPEND|
    os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
    log.SetOutput(file)
}

func clearScreen() {
    value, ok := clear[runtime.GOOS]
    if ok {
        value()
    } else {
        log.Println("[ERROR]: Platform not supported yet...")
        os.Exit(-1)
    }
}

func render() {
    oldSize := len(GetTodos())
    for  {
        // Check if a new TODO has been added and if 
        // required to do another poll
        updatedTodos := GetTodos()
        if newSize := len(updatedTodos); newSize != oldSize {
            clearScreen()
            for _, t := range updatedTodos {
                fmt.Printf("%s : %s\n", t.Project, t.Description)
            }
            oldSize = newSize
        }
    }
}

/*
* Declaring the clear functions for each os.
* Required to update the screen with new todos.
*/
func init() {
    clear = make(map[string]func())
    clear["linux"] = func() {
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func main() {
    setupFileLogging()
    http.HandleFunc("/",                HandlerInit)
    http.HandleFunc("/health-checker",  HandlerHealthCheck)
    http.HandleFunc("/todos",           HandlerTodos)
    http.HandleFunc("/addTodo",         HandlerAddTodo)
    log.Printf("[INFO]: listenning on port: %s", PORT)
    go render();
    log.Fatal(http.ListenAndServe(PORT, nil));
}
