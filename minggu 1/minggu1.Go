package main

import "fmt"

func main() {
    var x int

    fmt.Scanln(&x)
    if x == 2 {
        fmt.Println("dua")
    } else if x == 3 {
        fmt.Println("tiga")
    } else {
        fmt.Println("bebas")
    }
}