package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type score struct {
	id    string
	score int
}

func main() {
	scores := []score{}

	for {
		option := getOption()

		switch option {
		case "A", "a":
			scores = append(scores, addStudentScore())

		case "P", "p":
			printReport(scores)

		case "Q", "q":
			fmt.Println("Bye")
			os.Exit(123)
		}
	}
}

func getOption() string {
	fmt.Println("Please select an option:")
	fmt.Println()

	fmt.Println("A) Add a new score")
	fmt.Println("P) Print a report")
	fmt.Println("Q) Quit")
	fmt.Println()

	var option string
	fmt.Scanln(&option)
	fmt.Println()

	return option
}

func addStudentScore() score {
	fmt.Println("Enter student's ID and score:")

	var id, rawScore string
	fmt.Scanln(&id, &rawScore)
	s, _ := strconv.Atoi(rawScore) // Convert string to int

	fmt.Println()

	return score{id: id, score: s}
}

func printReport(scores []score) {
	fmt.Println("Student's scores")
	fmt.Println(strings.Repeat("-", 14))

	for _, s := range scores {
		fmt.Println(s.id, s.score)
	}

	fmt.Println()
}
