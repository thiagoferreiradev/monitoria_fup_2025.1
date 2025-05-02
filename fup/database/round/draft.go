package main

import "fmt"

func round(num float64) int {
    n := int(num)
    frac := num - float64(n)

    if frac >= 0.5 {
        return n + 1
    } else {
        return n
    }
}

func floor(num float64) int {
    return int(num)
}

func ceil(num float64) int {
    n := int(num) + 1

    return n
}

func main() {
    var c string
    var n float64

    fmt.Scan(&c, &n)

    switch c {
    case "r":
        fmt.Println(round(n))
    case "f":
        fmt.Println(floor(n))
    default:
        fmt.Println(ceil(n))
    }
}
