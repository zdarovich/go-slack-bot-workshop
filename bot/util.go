package bot

import (
	"fmt"
	"time"
)

func generateUserSequential() (string, string) {
	start := time.Now()

	name := getRandomName()
	pic := getRandomPictureUrl()

	printElapsed(start)

	return name, pic
}

func generateUserParallel() (string, string) {
	start := time.Now()

	nameChan := make(chan string)
	picChan := make(chan string)

	go func() {
		name := getRandomName()
		nameChan <- name
	}()

	go func() {
		pic := getRandomPictureUrl()
		picChan <- pic
	}()

	name := <-nameChan
	pic := <-picChan

	printElapsed(start)

	return name, pic
}

func printElapsed(start time.Time) {
	after := time.Since(start)
	fmt.Printf("time elapsed %s \n", after)
}
