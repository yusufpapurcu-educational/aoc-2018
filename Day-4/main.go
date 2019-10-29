package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var guards = make(map[int]guard)

type guard [][]int

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	var guardName, oldGuardName, day, start int
	for s.Scan() {
		var time, minute int
		text := s.Text()
		_, err := fmt.Sscanf(string(text[12:17]), "%d:%d", &time, &minute)
		if err != nil {
			log.Fatal(err)
		}
		switch string(text[19:24]) {
		case "Guard":
			fmt.Println("Guard Change")
			_, err := fmt.Sscanf(string(text[25:]), "#%d begins shift", &guardName)
			if err != nil {
				log.Fatal(err)
			}
			start = minute
			break
		case "falls":
			fmt.Println("falls")
			Awaken(minute-start-1, guardName, day)
			start = minute

			break
		case "wakes":
			fmt.Println("Wakes")
			start = minute

			break
		}
		if guardName != oldGuardName {
			if guards[guardName] == nil {
				day = 0
				guards[guardName][0] = make([]int, 60)
			} else {
				day = len(guards[guardName][day])
			}

			oldGuardName = guardName
		}
	}
	fmt.Println(guards[99][0])
}

func Awaken(take, guardName, day int) {
	list := guards[guardName][day]
	for i := 0; i <= take; i++ {
		list[i] = 1
	}
	fmt.Println(list)
	guards[guardName][day] = (list)
}
