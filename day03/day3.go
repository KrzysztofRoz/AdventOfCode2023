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
	//fmt.Println(SecondPart(path))
}

func SecondPart(filePath string) int {

	return 0
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
	//fmt.Println(symbols)
	//fmt.Println(engineParts)
	fmt.Println("Parts:", len(engineParts), " and Symbols: ", len(symbols))
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
		//fmt.Println(symbol.Show())
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
		//fmt.Println(enginePart.Show())
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
