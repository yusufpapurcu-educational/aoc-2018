package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type specs struct {
	guardName int
	day       int
	min       int
}
type guard struct {
	guardName int
	minute    int
}

var shafts = make(map[specs]int)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	var guardName, day, sleeptime = 0, 0, 0
	var days = make(map[int]int)
	for s.Scan() {
		var time, minute int
		text := s.Text()
		_, err := fmt.Sscanf(string(text[12:17]), "%d:%d", &time, &minute)
		if err != nil {
			log.Fatal(err)
		}
		switch string(text[19:24]) {
		case "Guard":
			days[guardName]++
			_, err := fmt.Sscanf(string(text[25:]), "#%d begins shift", &guardName)
			if err != nil {
				log.Fatal(err)
			}

			day = days[guardName]
			break
		case "falls":
			sleeptime = minute

			break
		case "wakes":

			fmt.Println("Wakes minute : ", minute-sleeptime)
			for i := sleeptime; i < minute; i++ {
				shafts[specs{guardName, day, i}] = 1
			}
			break
		}
	}
	// for i := 0; i < 60-laps; i++ {
	// 	shafts[specs{guardName, day, laps + i}] = 0
	// }
	days[guardName]++
	for i := range days {
		for c := 0; c < days[i]; c++ {
			fmt.Print(i, "  =  ", c, "     ")
			for b := 0; b < 60; b++ {
				if shafts[specs{i, c, b}] == 1 {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	}
	sleeper := make(map[int]int)
	sleepermin := make(map[guard]int)
	for i := range days {
		// I guard Name
		for c := 0; c < days[i]; c++ {
			// C day name
			for b := 0; b < 60; b++ {
				// B minute
				if shafts[specs{i, c, b}] == 1 {
					sleeper[i]++
					sleepermin[guard{i, b}]++
				}
			}
		}
	}
	mostSleeper := 0
	for i := range sleeper {
		if sleeper[i] > sleeper[mostSleeper] {
			mostSleeper = i
		}
	}
	fmt.Println(mostSleeper)
	mostSleepedMin := 0
	for i := 0; i < 60; i++ {
		if sleepermin[guard{mostSleeper, mostSleepedMin}] < sleepermin[guard{mostSleeper, i}] {
			mostSleepedMin = i
		}
	}
	fmt.Println(mostSleepedMin)
	fmt.Println(mostSleepedMin * mostSleeper)
}
