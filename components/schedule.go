package main

import "fmt"

func main() {
	people := []string{"Samuel", "Stanley", "Chidera", "Pius"}
	days := []string{"Mon", "Tue", "Wed", "Thu", "Fri"}

	weeks := 10

	// Create a 2D array to represent the schedule
	schedule := make([][]string, len(days))
	for i := range schedule {
		schedule[i] = make([]string, weeks)
	}

	// Fill in the schedule
	for j := 0; j < weeks; j++ {
		for i, day := range days {
			// Assign Chidera to Wednesdays
			if day == "Wed" {
				schedule[i][j] = "Chidera"
			} else {
				// Assign the other individuals randomly
				index := j%len(people)
				schedule[i][j] = people[index]
			}
		}
	}

	// Print the schedule
	for i, day := range days {
		fmt.Printf("%s:\t", day)
		for j := 0; j < weeks; j++ {
			fmt.Printf("%s\t", schedule[i][j])
		}
		fmt.Println()
	}
}