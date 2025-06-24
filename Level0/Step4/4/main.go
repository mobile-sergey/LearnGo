package main

import (
	"fmt"
    "time"
)

func main() {
    var t string
    fmt.Scanln(&t)
    parsed, _ := time.Parse("2006-01-02/15:04:05", t)
    result := fmt.Sprintf("Текущее время %d часов, %d минут. Ты точно не забыл про важные дела на сегодня?", parsed.Hour(), parsed.Minute())
    fmt.Println(result)
}