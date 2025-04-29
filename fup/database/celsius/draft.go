package main

import "fmt"

func main() {
    var celsius float64

    fmt.Scan(&celsius)

    fahreinheit := 1.8 * celsius + 32

    fmt.Printf("%.6f\n", fahreinheit)
}
