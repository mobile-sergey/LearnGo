package main

import (
	"fmt"
    "time"
    "math"
)

func main() {
    var date, firstname, lastname, middlename string
    var firstsum, secondsum, thirdsum float64
    fmt.Scanln(&date)
    fmt.Scanln(&firstname)
    fmt.Scanln(&lastname)
    fmt.Scanln(&middlename)
    fmt.Scanln(&firstsum)
    fmt.Scanln(&secondsum)
    fmt.Scanln(&thirdsum)
    newDate, _ := time.Parse("02.01.2006", date)
    newDate = newDate.AddDate(0, 0, 15)
    textDate := newDate.Format("02.01.2006")
    sum := firstsum + secondsum + thirdsum
    rub := int(math.Floor(sum))
    kop :=int(math.Round((sum - float64(rub)) * 100))
    result := fmt.Sprintf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\nДата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\nОбщая сумма выплат составит %d руб. %d коп.\n\nС уважением,\nГл. бух. Иванов А.Е.", lastname, firstname, middlename, textDate,rub, kop)
    fmt.Println(result)
}