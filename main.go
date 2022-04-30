package main

import (
	"encoding/csv"
	"flag"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := parseCsv()
	correctAnswers, totalQuestions := runQuiz(lines)
	Printf("You scored: %v/%v\n", correctAnswers, totalQuestions)
}

func parseCsv() [][]string {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(Sprintf("Failed to open the csv file: %v", csvFileName))
		os.Exit(1)
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	return lines
}

func runQuiz(lines [][]string) (int, string) {
	problems := parseProblems(lines)
	var correctAnswers = 0
	var totalQuestions = strconv.FormatInt(int64(len(problems)), 10)
	for i, p := range problems {
		Printf("Problem # %v: %v = \n", i+1, p.question)
		var answer string
		scan, err := Scanf("%s\n", &answer)
		if err != nil {
			exit("Failed to parse input")
		}
		if answer == p.answer {
			correctAnswers++
		}
		_ = scan
	}
	return correctAnswers, totalQuestions
}

func parseProblems(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	Println(msg)
	os.Exit(1)
}

type problem struct {
	question string
	answer   string
}
