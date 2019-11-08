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
	min := 31323123132131232
	for i := 65; i <= 90; i++ {
		tempUnit := Unit
		last := true
		tempUnit = CleanMaster(tempUnit, byte(i))
		for last {
			last, tempUnit = Process(tempUnit)
		}
		if len(tempUnit) < min {
			min = len(tempUnit)
		}
		fmt.Println("Tryed : ", string(byte(i)), "\nResult : ", len(tempUnit))
	}
	fmt.Println(min)

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

func CleanMaster(Unit string, whoClear byte) string {
	newUnit := ""
	for i := 0; i < len(Unit); i++ {
		if Unit[i] == whoClear || Unit[i]+32 == whoClear || Unit[i]-32 == whoClear {

		} else {
			newUnit += string(Unit[i])
		}
	}
	return newUnit
}
