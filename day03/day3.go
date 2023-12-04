package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type EnginePart struct {
	line  int
	start int
	end   int
	value int
}

type Symbol struct {
	mark  string
	line  int
	index int
}

func (enginePart *EnginePart) Show() string {
	result := fmt.Sprint("line:", enginePart.line, ",start:", enginePart.start, ",end:", enginePart.end, ",value:", enginePart.value)
	return result
}
func (symbol *Symbol) Show() string {
	result := fmt.Sprint("symbol:", symbol.mark, ",line:", symbol.line, ",index:", symbol.index)
	return result
}

func main() {
	path := "/home/krzysztof/Repos/AdventOfCode2023/inputs/input03.txt"
	fmt.Println(FirstPart(path))
	fmt.Println(SecondPart(path))
}

func SecondPart(filePath string) int {

	totalGearRatio := 0
	engineParts := make([]EnginePart, 0)
	symbols := make([]Symbol, 0)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		engineParts = append(engineParts, parseEngineParts(scanner.Text(), i)...)
		symbols = append(symbols, parseSymbols(scanner.Text(), i)...)
		i++
	}

	for _, gear := range symbols {
		if gear.mark == "*" {
			totalGearRatio += calculateGearRatio(gear, engineParts)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return totalGearRatio
}

func FirstPart(filePath string) int {

	partNumbersSum := 0
	engineParts := make([]EnginePart, 0)
	symbols := make([]Symbol, 0)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		engineParts = append(engineParts, parseEngineParts(scanner.Text(), i)...)
		symbols = append(symbols, parseSymbols(scanner.Text(), i)...)
		i++
	}

	for _, part := range engineParts {
		if partMatchSymbol(part, symbols) {
			partNumbersSum += part.value
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return partNumbersSum
}

func parseSymbols(engineString string, line int) []Symbol {
	notNumber := regexp.MustCompile("[^0-9\\s\\.]+")
	symbols := make([]Symbol, 0)

	for _, match := range notNumber.FindAllStringIndex(engineString, -1) {
		var symbol *Symbol
		symbol = new(Symbol)
		symbol.line = line
		symbol.index = match[0]
		symbol.mark = engineString[match[0]:match[1]]
		symbols = append(symbols, *symbol)
	}

	return symbols
}
func parseEngineParts(engineString string, line int) []EnginePart {
	re := regexp.MustCompile("\\d+")
	var err error
	engineParts := make([]EnginePart, 0)

	for _, match := range re.FindAllStringIndex(engineString, -1) {
		var enginePart *EnginePart
		enginePart = new(EnginePart)
		enginePart.line = line
		enginePart.start = match[0]
		enginePart.end = match[1] - 1
		enginePart.value, err = strconv.Atoi(engineString[match[0]:match[1]])
		engineParts = append(engineParts, *enginePart)
	}
	if err != nil {
		log.Fatal(err)
	}
	return engineParts
}

func partMatchSymbol(engPart EnginePart, symbols []Symbol) bool {
	result := false

	for _, symbol := range symbols {
		if engPart.line-1 <= symbol.line && engPart.line+1 >= symbol.line {
			if engPart.start-1 <= symbol.index && engPart.end+1 >= symbol.index {
				return true
			}
		}
	}
	return result
}
func calculateGearRatio(gear Symbol, engPart []EnginePart) int {
	result := 0
	var element1 *EnginePart
	element1 = new(EnginePart)
	var element2 *EnginePart
	element2 = new(EnginePart)
	var element3 *EnginePart
	element3 = new(EnginePart)

	for _, eng := range engPart {
		//fmt.Println(gear.Show())
		//fmt.Println(eng.Show())
		if matchGear(gear, eng) {
			fmt.Println("Match")
			if element1.value == 0 {
				fmt.Println("Insert to first")
				element1.value = eng.value
				fmt.Println(element1)
			} else if element2.value == 0 {
				fmt.Println("Insert to second")
				element2.value = eng.value
				fmt.Println(element2)

			} else if element3.value == 0 {
				fmt.Println("Insert to third")

				element3.value = 1
				break
			}
		}
	}
	if element3.value == 0 && element2.value != 0 {
		result = element1.value * element2.value
		fmt.Println("First element: ", element1.value, " Second element: ", element2.value)
		fmt.Println("Add ", result)
	}

	return result
}

func matchGear(gear Symbol, engPart EnginePart) bool {
	result := false

	if gear.line-1 <= engPart.line && gear.line+1 >= engPart.line {
		if (engPart.start <= gear.index+1 && engPart.start >= gear.index-1) || (engPart.end <= gear.index+1 && engPart.end >= gear.index-1) {
			result = true
		}
	}
	return result
}
