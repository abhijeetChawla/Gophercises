package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	problems := getProblems("problems.csv")
	fmt.Println(problems)
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
