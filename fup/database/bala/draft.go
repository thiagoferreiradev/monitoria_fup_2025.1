package main

import (
    "fmt"
    "math"
)

func main() {
    var p1_x, p1_y, p2_x, p2_y float64

    fmt.Scan(&p1_x, &p1_y, &p2_x, &p2_y)

    distancia := math.Sqrt(math.Pow(p2_x - p1_x, 2) + math.Pow(p2_y - p1_y, 2))

    fmt.Printf("%.2f\n", distancia)
}
