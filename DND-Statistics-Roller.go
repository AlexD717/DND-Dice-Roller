package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const totalSimulationTimes = 1000000

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

func worker(workerRepeat int, results chan<- []int, wg *sync.WaitGroup) {
	defer wg.Done()

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

	results <- highestRoll
}

func main() {
	startTime := time.Now()

	numWorkers := runtime.NumCPU()
	workerRepeat := totalSimulationTimes/numWorkers
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("Utilizing %d CPU Threads\n", numWorkers)

	var wg sync.WaitGroup

	results := make(chan []int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(workerRepeat, results, &wg)
	}

	wg.Wait()
	close(results)
	
	highestScore := 0
	highestRoll := []int{}
	for result := range(results) {
		resultScore := sum(result)
		if (resultScore > highestScore) {
			highestRoll = result
			highestScore = resultScore
		}
	}

	fmt.Printf("Highest Stats Rolled: ")
	fmt.Println(highestRoll)

	endTime := time.Now()
	timeTake := endTime.Sub(startTime)

	fmt.Println("Total time taken: " + timeTake.String())
	
	fmt.Println("Press Enter to exit...")
	var input string
	fmt.Scanln(&input)
}
