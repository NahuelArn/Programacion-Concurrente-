package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 1; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("aa")
    say("bb")
}