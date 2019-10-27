package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var fab fabric

	s := bufio.NewScanner(f)
	for s.Scan() {
		var id, x, y, w, h int
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("ID : %d -- (X : %d Y : %d) -- (W : %d H : %d)\n", id, x, y, w, h)
		fab.addClaim(id, x, y, w, h)

	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	for i := range overflap {
		fmt.Println(i, overflap[i])
	}
	fab.countClaim()
}

type xy struct{ x, y int }
type fabric struct {
	m map[xy]int
}

var overflap = make(map[int]bool)

func (f *fabric) addClaim(id, x, y, w, h int) {
	if f.m == nil {
		f.m = make(map[xy]int)
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			f.m[xy{x + i, y + j}]++
			if f.m[xy{x + i, y + j}] > 1 {
				overflap[id] = false
			}
		}
	}
}

func (f *fabric) countClaim() {
	counter := 0
	for i, _ := range f.m {
		if f.m[i] > 1 {
			counter++
		}
	}
	fmt.Println(counter)
}
