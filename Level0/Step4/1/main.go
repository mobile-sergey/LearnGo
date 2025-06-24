package main

import (
	"fmt"
	"math"
)

func main() {
    var a, b, c int
    fmt.Scanln(&a)
    fmt.Scanln(&b)
    fmt.Scanln(&c)
    min1 := math.Min(float64(a), float64(b))
    min := math.Min(min1, float64(c))
    fmt.Println(min)
}