package main

import "fmt"
import "time"

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    time.Println("done")

    done <- true
}


func main() {
    done := make(chan bool, 1)
    go worker(done)

    <-done
}


