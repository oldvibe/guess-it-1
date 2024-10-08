package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	Y "cc/functions"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var numbers []float64

	for scanner.Scan() {
		num, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		numbers = append(numbers, num)
		predictNextRange(numbers)
	}
}

func predictNextRange(numbers []float64) {
	if len(numbers) < 2 {
		fmt.Println("1 1000") 
		return
	}

	avg := Y.Average(numbers)
	sd := Y.StandardDev(numbers)

	minRange := math.Max(1, avg-sd)
	maxRange := avg + sd

	
	lastNum := numbers[len(numbers)-1]
	if lastNum > avg {
		minRange += 10
		maxRange += 20
	} else {
		minRange -= 20
		maxRange -= 10
	}

	
	minRange = math.Max(1, math.Min(minRange, 1000))
	maxRange = math.Max(1, math.Min(maxRange, 1000))

	fmt.Printf("%d %d\n", int(minRange), int(maxRange))
}
