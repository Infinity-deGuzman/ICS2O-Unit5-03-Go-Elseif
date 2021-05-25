// Created by: Infinity de Guzman
// Created on: May 2021
//
// This program sees what movies you're allowed to watch at your age.

package main

import "fmt"

func main() {
	var userAge int

	fmt.Println("How old are you? (integers only)")
	fmt.Println()
	fmt.Print("Age: ")
	fmt.Scanln(&userAge)

	if userAge >= 17 {
		print("You can see R-rated films/movies alone :)")
	} else if userAge >= 13 {
		print("You can see PG-13 films/movies alone :)")
	} else if userAge >= 5 {
		print("You can see G or PG films/movies alone :D")
	} else {
		print("Uhhhh… you’re too young to watch most things.")
	}
}
