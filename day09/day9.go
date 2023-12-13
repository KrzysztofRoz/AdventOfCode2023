package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "/home/krzysztof/Repos/AdventOfCode2023/inputs/input09.txt"
	fmt.Println(FirstPart(path))
	//fmt.Println(SecondPart(path))
}

func FirstPart(filePath string) int {
	result := 0
	valuesHistory := make([][]int, 0)
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		valuesHistory = append(valuesHistory, parseRowToArray(scanner.Text()))
		result += InterpolateRightValue(GetDifferences(valuesHistory[len(valuesHistory)-1]))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func parseRowToArray(row string) []int {
	result := make([]int, 0)
	var err error
	var number int

	for _, val := range strings.Split(row, " ") {
		number, err = strconv.Atoi(val)
		result = append(result, number)
	}
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func GetDifferences(valueHistory []int) [][]int {
	allZeros := false
	history := make([][]int, 0)
	history = append(history, valueHistory)
	for !allZeros {
		allZeros = true
		temp := make([]int, 0)
		lastRecord := history[len(history)-1]
		for i := 0; i < len(lastRecord)-1; i++ {
			temp = append(temp, lastRecord[i+1]-lastRecord[i])
		}
		history = append(history, temp)

		for i := range temp {
			if temp[i] != 0 {
				allZeros = false
			}
		}
	}

	return history
}

func InterpolateRightValue(history [][]int) int {
	for i := len(history) - 1; i > 0; i-- {
		lastElement := history[i-1][len(history[i-1])-1]
		secondLast := history[i][len(history[i])-1]
		history[i-1] = append(history[i-1], lastElement+secondLast)
	}
	fmt.Println(history)

	return history[0][len(history[0])-1]
}

func SecondPart(filePath string) int {
	result := 0

	return result
}
