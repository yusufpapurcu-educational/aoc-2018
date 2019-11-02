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
	var guardName, day, start, laps = 0, 0, 0, 0
	var days = make(map[int]int)
	var isSleep bool
	for s.Scan() {
		var time, minute int
		text := s.Text()
		_, err := fmt.Sscanf(string(text[12:17]), "%d:%d", &time, &minute)
		if err != nil {
			log.Fatal(err)
		}
		switch string(text[19:24]) {
		case "Guard":
			fmt.Println(guardName, laps)
			if guardName != 0 {
				for i := 0; i < 60-laps; i++ {
					shafts[specs{guardName, day, laps + i}] = 0
				}
			}
			days[guardName]++
			_, err := fmt.Sscanf(string(text[25:]), "#%d begins shift", &guardName)
			if err != nil {
				log.Fatal(err)
			}

			day = days[guardName]
			laps = 0
			start = 0
			break
		case "falls":
			if !isSleep {
				if start > minute {
					minute += 60
				}
				fmt.Println("Waken minute : ", minute-start)
				isSleep = true
				laps += minute - start
				start = minute % 60
			}
			break
		case "wakes":
			if isSleep {
				if start > minute {
					minute += 60
				}
				fmt.Println("Sleep minute : ", minute-start)
				for i := 0; i < minute-start; i++ {
					shafts[specs{guardName, day, laps + i}] = 1
				}
				laps += minute - start
				start = minute % 60
				isSleep = false
			}
			break
		}
	}
	// for i := 0; i < 60-laps; i++ {
	// 	shafts[specs{guardName, day, laps + i}] = 0
	// }
	days[guardName]++
	// for i := range days {
	// 	for c := 0; c < days[i]; c++ {
	// 		fmt.Print(i, "  =  ", c, "     ")
	// 		for b := 0; b < 60; b++ {
	// 			if shafts[specs{i, c, b}] == 1 {
	// 				fmt.Print("#")
	// 			} else {
	// 				fmt.Print(".")
	// 			}
	// 		}
	// 		fmt.Println()
	// 	}
	// }
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
	var mostg, mostd int
	for i := range days {
		for b := 0; b < 60; b++ {
			if sleepermin[guard{i, b}] >= sleepermin[guard{mostg, mostd}] {
				mostg = i
				mostd = b
			}
		}
	}
	var mostSleeper, mostSleepedMin int
	for i := range sleeper {
		if sleeper[i] >= sleeper[mostSleeper] {
			mostSleeper = i
		}
	}
	for i := 0; i < 60; i++ {
		if sleepermin[guard{mostSleeper, mostSleepedMin}] <= sleepermin[guard{mostSleeper, i}] {
			mostSleepedMin = i
		}
	}
	// for i := 0; i < 60; i++ {
	// 	fmt.Println(mostSleeper, "-", i, " : ", sleepermin[guard{mostSleeper, i}])
	// }
	fmt.Println(mostSleeper, mostSleepedMin, mostSleepedMin*mostSleeper)
	fmt.Println(mostg, mostd, mostg*mostd)

}
