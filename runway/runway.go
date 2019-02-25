package main

import "fmt"

func explodeBin(i, l int) []int {
	bin := make([]int, l)
	for i > 0 {
		bin[l-1] = i & 1
		i >>= 1
		l -= 1
	}
	return bin
}

func generateAllRunways(n int) [][]int {
	runways := make([][]int, 1<<uint(n))
	for i := 0; i < (1 << uint(n)); i++ {
		runways = append(runways, explodeBin(i, n))
	}
	return runways
}

func canStopRecursive(runway []int, start int, speed int, iterations *int) bool {
	*iterations += 1
	if speed == 0 {
		return false
	}
	if start < 0 || start > len(runway)-1 {
		return false
	}
	if runway[start] == 1 {
		return false
	}
	if start == len(runway)-1 {
		return true
	}
	for _, speed := range []int{speed - 1, speed, speed + 1} {
		if canStop := canStopRecursive(runway, start+speed, speed, iterations); canStop {
			return true
		}
	}
	return false
}

func main() {
	for runwayLen := 1; runwayLen <= 35; runwayLen++ {
		worst := 0
		runway := make([]int, runwayLen)
		runway[runwayLen-1] = 1
		iteration := 0
		canStopRecursive(runway, 0, 2, &iteration)
		if iteration > worst {
			worst = iteration
		}
		fmt.Printf("%d, ", worst)
	}
}
