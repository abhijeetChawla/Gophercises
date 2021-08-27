package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	problems := getProblems("problems.csv")
	quiz(problems)
}

func quiz(pArr []problem) {
	reader := bufio.NewReader(os.Stdin)
	totalPoints := len(pArr)
	count := 0
	for i, p := range pArr {
		fmt.Printf("Problem %d: %s \n", i, p.Q)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("There was an error please continue with the quiz for now")
			continue
		}
		answer := strings.ToLower(strings.TrimSpace(text))
		if answer == p.A {
			count++
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
