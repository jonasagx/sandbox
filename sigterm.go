package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time" // or "runtime"
)

func CheckToKill() {
    c := make(chan os.Signal, 2)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    fmt.Println(c)
    go func() {
        <-c
        fmt.Println("cleanup")
        os.Exit(0)
    }()
}

func main() {
    CheckToKill()
    for {
        fmt.Println("sleeping...")
        time.Sleep(1 * time.Second) // or runtime.Gosched() or similar per @misterbee
    }
}
