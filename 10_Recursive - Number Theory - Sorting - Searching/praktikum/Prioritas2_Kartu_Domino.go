package main

import "fmt"

func playingDomino(card [][]int, deck []int) interface{} {
	var result []int

	for i := 0; i < len(card); i++ {
		if card[i][0] == deck[0] || card[i][1] == deck[0] || card[i][0] == deck[1] || card[i][1] == deck[1] {
			result = append(result, card[i]...)
			return result
		}
	}

	return "tutup kartu"
}

func main() {
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))
}
