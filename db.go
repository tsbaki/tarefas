package main

import (
    "errors"
    "log"
    "os"
)

type Todo struct {
    Project string
    Description string
}

var ts[]Todo

func AddTodo(t *Todo) error {
    log.Printf("[INFO]: Adding TODO: %s | %s",
    t.Project, t.Description)
    if exists(ts, t) {
        return errors.New("Specified TODO already exists.")
    } else {
        ts = append(ts, Todo{t.Project, t.Description})
    }
    return nil
}

func RemoveTodo(todo *Todo) error {
    if exists(ts, t) {
        // TODO: Do your removing here
        nil
    } else {
        return errors.New("Specified TODO does not exist.")
    }
}

func GetTodos() []Todo {
    return ts
}

func WriteToFile(db *os.File) error {
    // Save TODOS to file
    return nil
}

func exists(slice []Todo, t *Todo) (bool) {
    for _, item := range slice  {
        if item.Project == t.Project && 
        item.Description == t.Description {
            return true
        }
    }
    return false
}
