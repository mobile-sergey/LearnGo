package main

import (
    "fmt"
    "time"
)

func main() {
    var now, past string
    const format = "2006-01-02"
    fmt.Scanln(&now)
    fmt.Scanln(&past)
    time1, _ := time.Parse(format, now)
    time2, _ := time.Parse(format, past)
    result := fmt.Sprintf("%d year ago", time1.Year() - time2.Year())
    fmt.Printf(result)
}