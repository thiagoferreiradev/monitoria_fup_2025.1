package main

import "fmt"

func main() {
    var n1, n2 int

    fmt.Scan(&n1, &n2)

    quociente := n1 / n2
    resto := n1 % n2

    fmt.Printf("%d %d\n", quociente, resto)
}
