package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type Problem struct {
	Question string
	Answer string
}

func main() {
	csvFile, _ := os.Open("problems.csv")
	csvReader := csv.NewReader(bufio.NewReader(csvFile))

	timeLimit := flag.Int("timer", 30, "Time limit for answering questions in seconds")

	flag.Parse()

	var problems []Problem

	for {
		line, error := csvReader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		problems = append(problems, Problem{
			Question: line[0],
			Answer: line[1],
		})
	}

	score := 0

	timer := time.NewTimer(time.Second * time.Duration(*timeLimit))
	defer timer.Stop()

	go func() {
		<-timer.C
		c := color.New(color.FgHiRed)
		c.Printf("\nTime's up! You scored ... %d\n", score)
		os.Exit(3)
	}()

	for index, problem := range problems {
		c := color.New(color.FgCyan)
		c.Printf("Problem #%d: %s = ", index + 1, problem.Question)

		var answer string
		fmt.Scan(&answer)

		if strings.TrimSpace(answer) == problem.Answer {
			score ++
		}
	}

	scoreColor := color.FgHiRed
	if score > 10 {
		scoreColor = color.FgGreen
	}

	c := color.New(scoreColor)

	c.Println("You scored ... ", score)
}

