package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var Unit string
	s := bufio.NewScanner(f)
	for s.Scan() {
		Unit = s.Text()
	}
	last := true
	for last {
		last, Unit = Process(Unit)
	}
	fmt.Println(len(Unit))
}

func Process(Unit string) (bool, string) {
	isProcess := false
	newUnit := ""
	for i := 0; i < len(Unit); i++ {
		if i == len(Unit)-1 {
			if Unit[i-1]+32 == Unit[i] || Unit[i-1] == Unit[i]+32 {
				isProcess = true
				i++
			} else {
				newUnit += string(Unit[i])
			}
		} else if Unit[i]+32 == Unit[i+1] || Unit[i] == Unit[i+1]+32 {
			isProcess = true
			i++
		} else {
			newUnit += string(Unit[i])
		}
	}
	return isProcess, newUnit
}

// dabCBAcaDA
// dabCBAcaDA
