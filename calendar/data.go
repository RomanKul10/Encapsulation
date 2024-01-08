package calendar

import "errors"

type Date struct {
	year  int // пишемо з малої щоб його не можна було імпортувати в інший пакет
	month int
	day   int
}

// створюєємо функцію сеттер(тобто присвоєння) який буде працювати для типу даних Date
func (d *Date) SetYear(year int) error { // (d *Date) *- означає що ми працюємо з копієєю початкової дати
	// передаємо параметр від користувача (year int)
	if year < 1 {
		return errors.New("Invalid year")
	}
	d.year = year // присвоюємо копії нашого Year = нові дані year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid day")
	}
	d.day = day
	return nil
}

// створюєємо геттер(тобто відавання) який буде віддавати нам значення
func (d *Date) Day() int { // функція для цьоготипу даних (d *Date) повертаємо Day з типом даних int
	return d.day
}
