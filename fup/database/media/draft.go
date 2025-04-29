package main

import "fmt"

func main() {
    var n1, n2 int

    fmt.Scan(&n1, &n2)

    media := (float64(n1) + float64(n2)) / 2.0

    fmt.Printf("%.1f\n", media)
}
