package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: cli <your_name>")
        return
    }
    fmt.Printf("Hello, %s!\n", os.Args[1])
}
