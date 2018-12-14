package main

import (
	"../../pkg/greetings"
	"fmt"
)

func main() {
	name := "James"
	greetings.GoodDay(name)
	greetings.GoodNight(name)

	if greetings.IsAM() {
		fmt.Println("Good morning", name)
	} else if greetings.IsAfternoon() {
		fmt.Println("Good afternoon", name)
	} else if greetings.IsEvening() {
		fmt.Println("Good evening", name)
	} else {
		fmt.Println("Good night", name)
	}
}
