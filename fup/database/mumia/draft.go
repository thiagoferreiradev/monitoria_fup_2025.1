package main

import "fmt"

func main() {
    var nome, class string
    var idade int

    fmt.Scan(&nome, &idade)

    if idade < 12 {
        class = "crianca"
    } else if idade < 18 {
        class = "jovem"
    } else if idade < 65 {
        class = "adulto"
    } else if idade < 1000 {
        class = "idoso"
    } else {
        class = "mumia"
    }

    fmt.Printf("%s eh %s\n", nome, class)
}
