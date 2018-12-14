package greetings

import (
	"fmt"
	"time"
)

var now time.Time

func GoodDay(name string) {
	fmt.Println("Good Day ", name)
}

func GoodNight(name string) {
	fmt.Println("Good Night ", name)
}

func IsAM() bool {
	now = time.Now()
	hour, _, _ := now.Clock()
	if hour < 5 || hour > 8 {
		return false
	}
	return true
}

func IsAfternoon() bool {
	now = time.Now()
	hour, _, _ := now.Clock()
	if hour < 8 || hour > 14 {
		return false
	}
	return true
}

func IsEvening() bool {
	now = time.Now()
	hour, _, _ := now.Clock()
	if hour < 18 || hour > 20 {
		return false
	}
	return true
}
