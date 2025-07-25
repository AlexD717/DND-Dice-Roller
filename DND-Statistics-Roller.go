package main

import (
	"fmt"
	"math/rand"
	"time"
)

const workerRepeat = 1000000

func removeFromList(list []int, value int) []int {
	for i, v := range list {
		if (list[i] == v) {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list
}

func min(list []int) int {
	if (len(list) == 0) {
		return 0
	}
	smallestValue := list[0]
	for _, v := range list {
		if (v < smallestValue) {
			smallestValue = v
		}
	}
	return smallestValue
}

func sum(list []int) int {
	total := 0
	for _, v := range list {
		total += v
	}
	return total
}

func rollStat() int {
	rolls := []int{}
	for i := 0; i < 4; i++ {
		rolls = append(rolls, rand.Intn(6)+1)
	}
	rolls = removeFromList(rolls, min(rolls))
	return sum(rolls)
}

func rollStats() []int {
	stats := []int{}
	for i := 0; i < 6; i++ {
		stats = append(stats, rollStat())
	}
	return stats
}

func worker() []int {
	highestRoll := []int{}
	highestScore := 0

	for i :=0; i < workerRepeat; i++ {
		currentRoll := rollStats()
		currentRollScore := sum(currentRoll)
		if (sum(currentRoll) > highestScore) {
			highestScore = currentRollScore
			highestRoll = currentRoll
		}
	}

	return highestRoll
}

func main() {
	startTime := time.Now()

	fmt.Println(worker())

	endTime := time.Now()
	timeTake := endTime.Sub(startTime)

	fmt.Println("Total time taken: " + timeTake.String())
}
