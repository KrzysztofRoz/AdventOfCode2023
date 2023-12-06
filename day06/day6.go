package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	path := "/home/krzysztof/Repos/AdventOfCode2023/inputs/input06.txt"
	fmt.Println(FirstPart(path))
	fmt.Println(SecondPart(path))
}

func FirstPart(filePath string) int {
	result := 1
	times := make([]int, 0)
	distances := make([]int, 0)
	reNum := regexp.MustCompile("\\d+")
	var isTime bool
	var tmp int

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		isTime, err = regexp.MatchString("Time:", scanner.Text())
		if isTime {
			for _, val := range reNum.FindAllString(scanner.Text(), -1) {
				tmp, err = strconv.Atoi(val)
				times = append(times, tmp)
			}
		} else {
			for _, val := range reNum.FindAllString(scanner.Text(), -1) {
				tmp, err = strconv.Atoi(val)
				distances = append(distances, tmp)
			}
		}

	}
	fmt.Println(times)
	fmt.Println(distances)
	for i := range times {
		result *= CalculatePossibilities(times[i], distances[i])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func SecondPart(filePath string) int {
	result := 1
	var time int
	var distance int
	reNum := regexp.MustCompile("\\d+")
	var isTime bool
	var tmp int

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		isTime, err = regexp.MatchString("Time:", scanner.Text())
		noSpace := strings.ReplaceAll(scanner.Text(), " ", "")
		if isTime {
			for _, val := range reNum.FindAllString(noSpace, -1) {
				tmp, err = strconv.Atoi(val)
				time = tmp
			}
		} else {
			for _, val := range reNum.FindAllString(noSpace, -1) {
				tmp, err = strconv.Atoi(val)
				distance = tmp
			}
		}

	}
	fmt.Println(time)
	fmt.Println(distance)
	result = CalculatePossibilities(time, distance)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func CalculatePossibilities(time, distance int) int {
	result := 0
	for i := 0; i <= time; i++ {
		road := i * (time - i)
		if road > distance {
			result++
		}
	}
	return result
}
