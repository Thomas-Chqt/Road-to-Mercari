package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var englishWords = []string{"water", "sun", "tree", "house", "cat", "dog", "bird", "car", "book", "computer"}

	stdin := bufio.NewScanner(os.Stdin)
	timeout := time.After(30 * time.Second)

	var score = 0

	go func() {
		for {
			rand.Seed(time.Now().UnixNano())
			randomIndex := rand.Intn(len(englishWords))
			var word = englishWords[randomIndex]
			var input = ""
			for input != word {
				fmt.Println(word)
				fmt.Print("-> ")
				if stdin.Scan() {
					input = stdin.Text()
				} else { 
					os.Exit(1)
				}
				if input == word {
					score++
				} else {
					fmt.Println("Wrong ! try Again.")
				}
			}
			
		}
	}()

	<-timeout
	fmt.Printf("Time's up! Score : %d\n", score);
	os.Exit(0)
}