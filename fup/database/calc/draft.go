package main

import "fmt"

func calc(n1 int, n2 int, op string) int {
    switch op {
    case "+":
        return n1 + n2
    case "-":
        return n1 - n2
    case "/":
        return n1 / n2
    case "*":
        return n1 * n2
    default:
        return -1
    }
}

func main() {
    var n1, n2 int
    var op string

    fmt.Scan(&n1, &n2, &op)

    fmt.Println(calc(n1, n2, op))
}
