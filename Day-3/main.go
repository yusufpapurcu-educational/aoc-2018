package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type xy struct{ x, y int }
type fabric struct {
	m map[xy]int
}

type data struct {
	id, x, y, w, h int
}

var overflap = make(map[int]bool)

func main() {
	f, err := os.Open("input.txt") // Dosya acma
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // Islem bitince dosya kapama

	var fab fabric           // Haritamiz
	var datas []data         // Tum konum verilerinin tutulacagi dizi
	s := bufio.NewScanner(f) // Scanner ayarlaniyor
	for s.Scan() {
		var id, x, y, w, h int
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h) // Tum degerler aliniyor
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("ID : %d -- (X : %d Y : %d) -- (W : %d H : %d)\n", id, x, y, w, h)
		overflap[id] = true                         // Tum idler true yapiliyor. Gorevin ikinci kismi icin onemli
		fab.addClaim(id, x, y, w, h)                // Harita ciziliyor
		datas = append(datas, data{id, x, y, w, h}) // data dizisi guncelleniyor
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	fab.countClaim()      // Kesismeler sayiliyor
	fab.isOverflap(datas) // Hic kesismeyen id bulunuyor
}

// Bu method harita cizmek icin
func (f *fabric) addClaim(id, x, y, w, h int) {
	if f.m == nil { // Harita bos ise olusturuluyor
		f.m = make(map[xy]int)
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			f.m[xy{x + i, y + j}]++ // Cizme islemi karenin degeri bir arttiliyor kesisim bulmak icin ideal
		}
	}
}

// Bu method kesisim sayma icin. Aciklama geregi duymuyorum
func (f *fabric) countClaim() {
	counter := 0
	for i := range f.m {
		if f.m[i] > 1 {
			counter++
		}
	}
	fmt.Println(counter)
}

// Bu method kesismeyenleri bulmak icin
func (f *fabric) isOverflap(datas []data) {
	for _, data := range datas {
		for i := 0; i < data.w; i++ {
			for j := 0; j < data.h; j++ {
				if f.m[xy{data.x + i, data.y + j}] > 1 { // Eger kesisim var ise overflap dizisindeki degerini false yaparak donguden atiyor.
					overflap[data.id] = false
					break
				}
			}
			if !overflap[data.id] {
				break
			}
		}
	}
	for i := range overflap {
		if overflap[i] {
			fmt.Println(i)
		}
	}
}
