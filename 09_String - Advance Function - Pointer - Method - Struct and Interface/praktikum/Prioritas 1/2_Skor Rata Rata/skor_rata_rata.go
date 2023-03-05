package main

import (
	"fmt"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var total int
	for _, score := range s.score {
		total += score
	}
	return float64(total) / float64(len(s.score))
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]
	for i, score := range s.score {
		if score < min {
			min = score
			name = s.name[i]
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]
	for i, score := range s.score {
		if score > max {
			max = score
			name = s.name[i]
		}
	}
	return max, name
}

func takeInputFromUser(a *Student) {
	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Urut ", i+1, " Nama Siswa : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)

		var score int
		fmt.Print("Masukkan ", name, " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}
}

func main() {
	var a = Student{name: []string{}, score: []int{}}

	takeInputFromUser(&a)

	fmt.Println("\n\nNilai rata-rata siswa ", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Nilai siswa maximum :", nameMax, "(", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Nilai siswa minimum :", nameMin, "(", scoreMin, ")")
}
