package main

import "fmt"

func main() {
    var qtd_p1, qtd_p2, qtd_p3 int
    var val_p1, val_p2, val_p3, dinheiro float64

    fmt.Scan(&qtd_p1, &qtd_p2, &qtd_p3,
                &val_p1, &val_p2, &val_p3,
                &dinheiro)

    total := float64(qtd_p1) * val_p1 +
                float64(qtd_p2) * val_p2 +
                float64(qtd_p3) * val_p3

    troco := dinheiro - total

    fmt.Printf("%.2f\n", troco)
}
