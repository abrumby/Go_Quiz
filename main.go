package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Answer   string
}

type Quiz struct {
	Problems []Problem
	Score    int
}

func main() {
	var file string
	var quizTime int
	var randomize bool
	flag.StringVar(&file, "file", "problems.csv", "--file=path/to/problems/file")
	flag.IntVar(&quizTime, "timeout", 30, "--time=15")
	flag.BoolVar(&randomize, "randomize", false, "--randomize=true")
	flag.Parse()

	AwaitStart(quizTime)
	quiz := Quiz{
		Problems: ParseProblemsFrom(file, randomize),
		Score:    0,
	}
	go func() {
		<-time.After(time.Duration(quizTime) * time.Second)
		TimeUpMessage()
		OutputMessage(quiz.Score, len(quiz.Problems))
		os.Exit(0)
	}()
	RunQuiz(&quiz)
	OutputMessage(quiz.Score, len(quiz.Problems))
}

func AwaitStart(quizTime int) {
	fmt.Printf("You have %v seconds to finish the quiz!\n", quizTime)
	fmt.Print("Press 'Enter' to start the quiz")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func ParseProblemsFrom(pathToFile string, randomize bool) []Problem {
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal("File does not exists")
	}
	reader := csv.NewReader(bufio.NewReader(file))
	var problems []Problem
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		problems = append(problems, Problem{
			Question: line[0],
			Answer:   line[1],
		})
	}
	if randomize {
		//randomize using current machine time
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(problems), func(i, j int) {
			problems[i], problems[j] = problems[j], problems[i]
		})
	}
	return problems
}
func RunQuiz(q *Quiz) {
	reader := bufio.NewReader(os.Stdin)
	for _, problem := range q.Problems {
		AskQuestion(&problem)
		answer := ReadLine(reader)
		if problem.Answer == answer {
			q.Score++
		}
	}
}

func AskQuestion(p *Problem) {
	fmt.Print(p.Question + " = ")
}

func ReadLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func TimeUpMessage() {
	fmt.Println("\rTime is up!")
}

func OutputMessage(correctAnswersCount int, problemsCount int) {
	fmt.Printf("\rYou scored %v out of %v!\n", correctAnswersCount, problemsCount)
}
