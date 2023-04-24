package main

import (
	"fmt"
	"strings"
)

func countLetters(text string) <-chan map[rune]int {
	letterCount := make(chan map[rune]int)

	go func() {
		count := make(map[rune]int)
		for _, letter := range text {
			count[letter]++
		}
		letterCount <- count
	}()

	return letterCount
}

func mergeCountChannels(channels ...<-chan map[rune]int) <-chan map[rune]int {
	merged := make(chan map[rune]int)

	go func() {
		eachChannelCount := make(map[rune]int)
		for _, channel := range channels {
			for letter, count := range <-channel {
				eachChannelCount[letter] += count
			}
		}
		merged <- eachChannelCount
	}()

	return merged
}

func main() {
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed non risus. Suspendisse lectus tortor, dignissim sit amet, adipiscing nec, ultricies sed, dolor. Cras elementum ultrices diam. Maecenas ligula massa, varius a, semper congue, euismod non, mi. Proin porttitor, orci nec nonummy molestie, enim est eleifend mi, non fermentum diam nisl sit amet erat. Duis semper. Duis arcu massa, scelerisque vitae, consequat in, pretium a, enim. Pellentesque congue. Ut in risus volutpat libero pharetra tempor. Cras vestibulum bibendum augue. Praesent egestas leo in pede. Praesent blandit odio eu enim. Pellentesque sed dui ut augue blandit sodales. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Aliquam nibh. Mauris ac mauris sed pede pellentesque fermentum. Maecenas adipiscing ante non diam sodales hendrerit."

	text = strings.ToLower(text)
	sentences := strings.Split(text, ".")

	var channels []<-chan map[rune]int
	for _, sentence := range sentences {
		channels = append(channels, countLetters(sentence))
	}

	merged := mergeCountChannels(channels...)

	for letter, count := range <-merged {
		fmt.Printf("%s: %d\n", string(letter), count)
	}
}
