package main

import (
	"bufio"
	"fmt"
	"os"
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
	// part1 := solvePart1(input)
	part2 := solvePart2(input)

	// fmt.Println("Part 1:", part1)
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
	safeUnsafe := []int{}

	// strings to int
	for lineIndex, line := range input {
		fmt.Printf("Processing line %d: %s\n", lineIndex, line)
		s := strings.Split(line, " ")

		isSafe := true
		var isPositive bool

		for i := 1; i < len(s); i++ {
			curr, _ := strconv.Atoi(s[i])
			prev, _ := strconv.Atoi(s[i-1])
			diff := curr - prev

			fmt.Printf("  Comparing s[%d]=%d and s[%d]=%d, diff=%d\n", i-1, prev, i, curr, diff)

			// not the same int
			if curr == prev {
				fmt.Println("    Found identical numbers, marking as unsafe")
				isSafe = false
				break
			}

			// difference must be between -3 and 3
			if diff < -3 || diff > 3 {
				fmt.Println("    Difference out of range, marking as unsafe")
				isSafe = false
				break
			} else {
				if i > 1 {
					currentPositive := diff > 0
					if currentPositive != isPositive {
						fmt.Println("    Direction of difference changed, marking as unsafe")
						isSafe = false
						break
					}
				} else {
					isPositive = diff > 0
				}
			}
		}

		if isSafe {
			fmt.Println("  Line is safe")
			safeUnsafe = append(safeUnsafe, 1)
		} else {
			fmt.Println("  Line is unsafe")
			safeUnsafe = append(safeUnsafe, 0)
		}
	}

	result := 0
	for _, v := range safeUnsafe {
		result += v
	}

	fmt.Printf("Safe count: %d\n", result)
	return result
}

func solvePart2(input []string) int {
	// wrong solution!! might do again in .py or use another approach and more time
	safeUnsafe := []int{}

	for lineIndex, line := range input {
		fmt.Printf("Processing line %d: %s\n", lineIndex, line)
		s := strings.Split(line, " ")
		issues := 0

		lastGoodIndex := 0 // Index of the last valid element

		for i := 1; i < len(s); i++ {
			curr, err1 := strconv.Atoi(s[lastGoodIndex])
			next, err2 := strconv.Atoi(s[i])
			if err1 != nil || err2 != nil {
				fmt.Println("    Error converting to int, skipping")
				issues++
				if issues > 1 {
					break
				}
				continue
			}

			diff := next - curr
			fmt.Printf("  Comparing s[%d]=%d and s[%d]=%d, diff=%d\n", lastGoodIndex, curr, i, next, diff)

			// Check for identical numbers
			if next == curr {
				fmt.Println("    Found identical numbers, incrementing issues and skipping")
				issues++
				if issues > 1 {
					break
				}
				continue
			}

			// Check if the difference is out of range
			if diff < -3 || diff > 3 {
				fmt.Println("    Difference out of range, incrementing issues and skipping")
				issues++
				if issues > 1 {
					break
				}
				continue
			}

			// Check if direction of difference changes (consistent with previous)
			if lastGoodIndex > 0 {
				prev, err := strconv.Atoi(s[lastGoodIndex-1])
				if err != nil {
					fmt.Println("    Error converting previous to int, incrementing issues and skipping")
					issues++
					if issues > 1 {
						break
					}
					continue
				}
				// Determine the direction of the current difference
				currentPositive := diff > 0
				// Determine the direction of the previous difference
				prevDiff := curr - prev
				prevPositive := prevDiff > 0

				if currentPositive != prevPositive {
					fmt.Println("    Direction of difference changed, incrementing issues and skipping")
					issues++
					if issues > 1 {
						break
					}
					continue
				}
			}

			// If no issues, update lastGoodIndex
			lastGoodIndex = i
		}

		fmt.Printf("  Issues for line %d: %d\n", lineIndex, issues)

		// Mark line as safe or unsafe
		if issues > 1 {
			fmt.Println("  Line is unsafe")
			safeUnsafe = append(safeUnsafe, 0)
		} else {
			fmt.Println("  Line is safe")
			safeUnsafe = append(safeUnsafe, 1)
		}
	}

	// Calculate the final result
	result := 0
	for _, v := range safeUnsafe {
		fmt.Printf("Safe/Unsafe: %d\n", v)
		result += v
	}

	finalResult := result
	fmt.Printf("Final result: %d\n", finalResult)
	return finalResult
}
