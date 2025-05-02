package main

import "fmt"

func main() {
    var n1, n2 int

    fmt.Scan(&n1, &n2)

    diferenca := n1 - n2

    if diferenca < 0 {
        diferenca = diferenca * -1
    }

    fmt.Println(diferenca)
}
