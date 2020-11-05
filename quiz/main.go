package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func main() {

	csvFile := flag.String("csv", "problems.csv", "enter csv filename")
	flag.Parse()

	file, err := os.Open(*csvFile)

	if err != nil {
		exit(fmt.Sprintf("Not able to open the file %v", *csvFile))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		exit("Error in parsing csv")
	}

	problems := parseLines(lines)
	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem %d : %s \n", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == p.answer {
			correct++
		}
	}

	fmt.Printf("You scored %d outof %d\n", correct, len(problems))
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
