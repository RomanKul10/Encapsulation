package main

// Інкапсуляція - метод приховання частини коду від іншої частини коду за допомогою методів сеттер(присвоєння) і геттер(віддавання)
import (
	"fmt"
	"log"
	"myapp/calendar"
)

func main() {
	dateUsers := calendar.Date{} // створюємо зміну dateUsers і імпортуємо нашу дату з calendar.Date
	err := dateUsers.SetYear(3110)
	if err != nil { //якщо наша помилка не дорівнюєє нашому значенню
		log.Fatal(err) // то показуємо нашу помилку і закриваємо програму
	}

	fmt.Println(dateUsers)
	fmt.Println(dateUsers.Day()) // вивести конкретне поле, у даному випадку "день"

}
