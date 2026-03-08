package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup

func hello() {
    defer wg.Done()
    fmt.Println("Hello from Go!")
}

func main() {
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go hello()
    }
    wg.Wait()
}