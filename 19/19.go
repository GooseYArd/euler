package main
import (
	"fmt"
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

func century_anchor(year int) int {

	var dayidx = map[string] int {
		"sunday" : 0,
		"monday":  1,
		"tuesday": 2,
		"wednesday": 3,
		"thursday": 4,
		"friday": 5,
		"saturday": 6, 
	}

	c := (year / 100) + 1
	return (((5 * c + ((c-1) / 4)) % 7) + dayidx["thursday"]) % 7
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

	// days := []string{
	//  	"sunday", 
	//  	"monday", 
	//  	"tuesday", 
	//  	"wednesday", 
	//  	"thursday", 
	//  	"friday", 
	//  	"saturday",
	// }
	
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

	sundays := 0
	for i:= 1901; i <= 2000; i++ {
		anchor := century_anchor(i)		
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
		for month := range months {	
			known_dday := dday_mon(i, month)
			for ; known_dday > 7;  {
				known_dday -= 7
			}	
			first := doomsday - (known_dday - 1)
			if first == 0 {
				sundays++
			}
			//if first < 0 {
			//	first += 7
			//}
			//fmt.Printf("If the %v of %v is a %v then the 1st is a %v\n", known_dday, months[month], days[doomsday], days[first])
		}
	}
	
	fmt.Printf("%v sundays fell on the first of a month.\n", sundays)
	
}




 