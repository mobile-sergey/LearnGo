package main

import "fmt"

func main() {
    var name string
    fmt.Scanln(&name)
	var flat int
    fmt.Scanln(&flat)
    var secret int
    fmt.Scanln(&secret)
    var time float64
    fmt.Scanln(&time)
    result := fmt.Sprintf("Привет, %s! Приглашаю тебя на соревнование по программированию, которое пройдёт, как всегда, в квартире %d. Оно будет длиться примерно %.1f часа. Не забудь секретный пароль для входа: %d.", name, flat, time, secret)
    fmt.Println(result)
}