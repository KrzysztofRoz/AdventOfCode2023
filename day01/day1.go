package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	path := "../inputs/input01.txt"
	result01 := FirstPart(path)
	fmt.Println(result01)
	result02 := SecondPart(path)
	fmt.Println(result02)
}

func SecondPart(filePath string) int {
	dictionary := make(map[string]int)
	dictionary["one"] = 1
	dictionary["two"] = 2
	dictionary["three"] = 3
	dictionary["four"] = 4
	dictionary["five"] = 5
	dictionary["six"] = 6
	dictionary["seven"] = 7
	dictionary["eight"] = 8
	dictionary["nine"] = 9

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	calibrationValues := []int{}

	reNumb := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine)")
	re := regexp.MustCompile("([0-9]{1}|one|two|three|four|five|six|seven|eight|nine)")

	for scanner.Scan() {

		canonString := reNumb.ReplaceAllStringFunc(scanner.Text(), func(match string) string {
			return repeatLastLetter(match)
		})
		digitsInLine := re.FindAllString(canonString, -1)

		for i := range digitsInLine {
			val, ok := dictionary[digitsInLine[i]]
			if ok {
				digitsInLine[i] = strconv.Itoa(val)
			}
		}

		first, err := strconv.Atoi(digitsInLine[0])
		last, err := strconv.Atoi(digitsInLine[len(digitsInLine)-1])
		if err != nil {
			log.Fatal(err)
		}
		calibrationValues = append(calibrationValues, (10*first + last))
	}
	result := 0
	for i := range calibrationValues {
		result += calibrationValues[i]

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func repeatLastLetter(s string) string {
	if len(s) > 0 {
		lastLetter := s[len(s)-1:]
		return s[:len(s)-1] + lastLetter + lastLetter
	}
	return s
}

func FirstPart(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	calibrationValues := []int{}

	re := regexp.MustCompile("[0-9]{1}")

	for scanner.Scan() {
		digitsInLine := re.FindAllString(scanner.Text(), -1)
		first, err := strconv.Atoi(digitsInLine[0])
		last, err := strconv.Atoi(digitsInLine[len(digitsInLine)-1])
		if err != nil {
			log.Fatal(err)
		}
		calibrationValues = append(calibrationValues, (10*first + last))
	}

	result := 0
	for i := range calibrationValues {
		result += calibrationValues[i]

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
