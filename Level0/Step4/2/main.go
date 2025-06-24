package main

import (
    "fmt"
    "math"
)

func main() {
    var x1, x2 float64
    fmt.Scanln(&x1)
    fmt.Scanln(&x2)
    result := math.Round(math.Max(x1, x2))
    fmt.Println(result)
}