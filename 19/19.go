package main
import (
	"fmt"
//	"math"
)

func isleap(year int) bool {
	if year % 4 == 0 {
		if year % 100 == 0 {
			if year % 400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

func dday_mon(year int, month int) int {
	ddays := []int{ 3, 28, 7, 4, 9, 6, 11, 8, 5, 10, 7, 12 }
	if isleap(year) {
		ddays[0]++
		ddays[1]++
	}
	return ddays[month]
}

func main() {

	// 
	// 
	// You are given the following information, but you may prefer to do some research for yourself.
	// 
	// 1 Jan 1900 was a Monday.
	// Thirty days has September,
	// April, June and November.
	// All the rest have thirty-one,
	// Saving February alone,
	// Which has twenty-eight, rain or shine.
	// And on leap years, twenty-nine.
	// A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
	// 
	// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
	// 
	// 

	days := []string{
		"sunday", 
		"monday", 
		"tuesday", 
		"wednesday", 
		"thursday", 
		"friday", 
		"saturday",
	}
	
	months := []string{
		"january",
		"february",
		"march",
		"april",
		"may",
		"june",
		"july",
		"august",
		"september",
		"october",
		"november",
		"december"}

	var dayidx = map[string] int {
		"sunday" : 0,
		"monday":  1,
		"tuesday": 2,
		"wednesday": 3,
		"thursday": 4,
		"friday": 5,
		"saturday": 6, 
	}
	
	anchor := dayidx["wednesday"]
	fmt.Printf("anchor for 20th century: %v\n", days[anchor])
	
	for i:= 1900; i < 1901; i++ {
		t := i % 100
		if t % 2 == 1 {
			t += 11
		}
		t = t / 2
		if t % 2 == 1 {
			t += 11
		}
		t = 7 - (t % 7)
		doomsday := (anchor + t) % 7
		
		fmt.Printf("doomsday for year %v is %v\n", i, days[doomsday])
                   
		for month := range months {
			dday := dday_mon(i, month)
			//fmt.Printf("%v/1/%v was a %v\n", month+1, i, days[(doomsday + offset) % 7])
			//fmt.Printf("the %vth of %v is a %v\n", dday, months[month], days[doomsday])
			fmt.Printf("the 1st of %v is a %v, %v\n",  months[month], ((doomsday % 7) + 1), dday)
			//fmt.Printf("doomsday - offset: %v\n", doomsday - offset)
		}
	}
	


	
}




 