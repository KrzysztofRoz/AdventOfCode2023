package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	path := "/home/krzysztof/Repos/AdventOfCode2023/inputs/test02.txt"
	fmt.Println(FirstPart(path))
	fmt.Println(SecondPart(path))
}

func FirstPart(filePath string) int {
	result := 0
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func SecondPart(filePath string) int {
	result := 0

	return result
}
