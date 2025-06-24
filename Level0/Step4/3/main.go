package main

import (
    "fmt"
    "math"
)

func main() {
    var x, y, z, m, n float64
    fmt.Scanln(&x)
    fmt.Scanln(&y)
    fmt.Scanln(&z)
    fmt.Scanln(&m)
    fmt.Scanln(&n)
    result := (5.0 * x * math.Sin(2.0 * y)) / (z + math.Pow(n, math.Log(m)))
    fmt.Println(result)
}