package main

import "fmt"

func main() {
    var p, a, b int

    fmt.Scan(&p, &a, &b)

    if (p == 0 && ((a + b) % 2) == 0) || (p == 1 && ((a + b) % 2) != 0) {
        fmt.Println(0)
    } else {
        fmt.Println(1)
    }
}
