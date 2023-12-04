package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ScrachCard struct {
	Id          int
	winNumbers  []int
	cardNumbers []int
}

func main() {
	path := "/home/krzysztof/Repos/AdventOfCode2023/inputs/input04.txt"
	fmt.Println(FirstPart(path))
}

func FirstPart(filePath string) int {
	result := 0
	counter := 0
	cards := make([]ScrachCard, 0)
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		cards = append(cards, ParseScrachCard(scanner.Text()))
	}

	for _, card := range cards {
		counter = 0
		for _, num := range card.winNumbers {
			if Contains(card.cardNumbers, num) {
				counter++
			}
		}
		if counter != 0 {
			result += int(math.Pow(2, float64(counter-1)))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result

}

func ParseScrachCard(cardString string) ScrachCard {
	var card *ScrachCard
	card = new(ScrachCard)
	var err error
	var cardParts []string
	var lotery []string
	var temp int

	cardParts = strings.Split(cardString, ": ")
	reg := regexp.MustCompile("(\\d+)")
	card.Id, err = strconv.Atoi(reg.FindAllString(cardParts[0], -1)[0])
	fmt.Println(card.Id)
	lotery = strings.Split(cardParts[1], " | ")
	for _, val := range reg.FindAllString(lotery[0], -1) {
		temp, err = strconv.Atoi(val)
		card.winNumbers = append(card.winNumbers, temp)
	}
	for _, val := range reg.FindAllString(lotery[1], -1) {
		temp, err = strconv.Atoi(val)
		card.cardNumbers = append(card.cardNumbers, temp)
	}
	fmt.Println(card)
	if err != nil {
		log.Fatal(err)
	}

	return *card
}

func Contains(slice []int, value int) bool {
	result := false

	for i := range slice {
		if slice[i] == value {
			result = true
			break
		}
	}

	return result
}
