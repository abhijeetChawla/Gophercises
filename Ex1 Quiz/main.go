package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "File to be used. it should in format of question,answer")
	timelimit := flag.Int("time", 30, "Time limit for the quiz in seconds")
	flag.Parse()
	problems := getProblems(*csvFile)
	quiz(problems, *timelimit)
}

func quiz(pArr []problem, timeLimit int) {
	reader := bufio.NewReader(os.Stdin)
	totalPoints := len(pArr)
	count := 0

	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
problemLoop:
	for i, p := range pArr {
		fmt.Printf("Problem %d: %s \n", i, p.Q)
		answerCh := make(chan string)
		go func() {
			ans, _ := reader.ReadString('\n')
			answerCh <- ans
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			answer = strings.ToLower(strings.TrimSpace(answer))
			if answer == p.A {
				count++
			}
		}
	}

	if count >= totalPoints/2 {
		fmt.Print("Congratulations! ")
	}
	fmt.Printf("You scored %d out of %d", count, totalPoints)
}

type problem struct {
	Q string
	A string
}

func getProblems(fileName string) []problem {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(f)
	csvLines, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	var ret []problem
	for _, lines := range csvLines {
		q := strings.TrimSpace(lines[0])
		a := strings.TrimSpace(lines[1])
		ret = append(ret, problem{q, a})
	}
	return ret
}
