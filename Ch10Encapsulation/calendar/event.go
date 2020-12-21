package calendar

import "errors"


type Event struct {
	title string
	Date
}
func (e *Event) SetTitle(title string) error{
	if len(title) > 30{
		return errors.New("invalid Title: too long")
	}
	e.title = title
	return nil
}
func (e *Event) Title() string {
	return e.title
}


type Date struct {
	year  int
	month int
	day   int
}
func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("invalid Year")
	}
	d.year = year
	return nil
}
func (d *Date) Year() int {
	return d.year
}


func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}
func (d *Date) Month() int {
	return d.month
}


func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid day")
	}
	d.day = day
	return nil
}
func (d *Date) Day() int {
	return d.day
}