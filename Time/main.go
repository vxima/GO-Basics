package main

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	const layout = "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now()

	passed, _ := time.Parse("January 2, 2006 15:04:05", date)
	return passed.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	ourDate, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour := ourDate.Hour()
	if hour >= 12 && hour < 18 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	data := Schedule(date)
	dia_semana := data.Weekday()
	dia := data.Day()
	mes := data.Month()
	ano := data.Year()
	hora := data.Hour()
	min := data.Minute()
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", dia_semana, mes, dia, ano, hora, min)

}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()
	anniversary := fmt.Sprintf("%d-09-15", currentYear)
	t, _ := time.Parse("2006-01-02", anniversary)
	return t
}

func main() {
	//dateString := "7/25/2019 13:45:00"
	//fmt.Println(Schedule(dateString))
	fmt.Println(HasPassed("December 9, 2112 11:59:59"))
	fmt.Println(IsAfternoonAppointment("Friday, March 8, 1974 12:02:02"))
	//fmt.Println(Description(dateString))
	//fmt.Println(AnniversaryDate())
}
