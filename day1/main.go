package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	left  int
	right int
}

func main() {
	input, err := readInput("./input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Implement your solution logic here
	part1 := solvePart1(input)
	part2 := solvePart2(input)

	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func solvePart1(input []string) int {

	left := []int{}
	right := []int{}

	// strings to int
	for _, line := range input {
		s := strings.Split(line, "   ")
		leftInt, _ := strconv.Atoi(s[0])
		rightInt, _ := strconv.Atoi(s[1])
		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	// order in place
	sort.Sort(sort.IntSlice(left))
	sort.Sort(sort.IntSlice(right))

	pairs := make([]Pair, len(left))
	for i := range left {
		pairs[i] = Pair{left[i], right[i]}
	}

	// Pair up the smallest number in the left list with the smallest number in the right list
	// figure out how far apart the two numbers are; you'll need to add up all of those distances

	distance := 0
	for i := range left {
		// never negative, calculate absolute value
		distance += int(math.Abs(float64(pairs[i].right - pairs[i].left)))
	}

	return distance
}

func solvePart2(input []string) int {
	left := []int{}
	right := []int{}

	// strings to int
	for _, line := range input {
		s := strings.Split(line, "   ")
		leftInt, _ := strconv.Atoi(s[0])
		rightInt, _ := strconv.Atoi(s[1])
		left = append(left, leftInt)
		right = append(right, rightInt)
	}

	similarityLeftRight := make([]int, len(left))

	pairs := make([]Pair, len(left))
	for i := range left {
		pairs[i] = Pair{left[i], right[i]}
	}

	for i := range pairs {
		similarityLeftRight[i] = 0
		// Überprüfen, wie oft der linke Wert in der rechten Liste vorkommt
		for j := range pairs {
			if pairs[i].left == pairs[j].right {
				similarityLeftRight[i] += 1
			}
		}
	}

	finalScore := 0
	pairs = make([]Pair, len(left))
	for i := range left {
		pairs[i] = Pair{left[i], similarityLeftRight[i]}
	}

	for i := range left {
		fmt.Print(pairs[i].left, " ", pairs[i].right, "\n")
		finalScore += pairs[i].left * pairs[i].right
	}

	return finalScore
}
