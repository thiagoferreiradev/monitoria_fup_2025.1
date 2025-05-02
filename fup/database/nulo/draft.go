package main

import "fmt"

func main() {
    var n int

    fmt.Scan(&n)

    if n > 0 {
        fmt.Println("+")
    } else if n < 0 {
        fmt.Println("-")
    } else {
        fmt.Println("0")
    }
}
