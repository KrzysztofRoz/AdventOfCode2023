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

const (
	MaxRed   = 12
	MaxGreen = 13
	MaxBlue  = 14
)

type Game struct {
	Id    int
	Red   int
	Blue  int
	Green int
}

func main() {
	path := "/home/krzysztof/Repos/AdventOfCode2023/inputs/input02.txt"
	fmt.Println(FirstPart(path))
	fmt.Println(SecondPart(path))
}

func FirstPart(filePath string) int {
	validIdCount := 0

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	games := make([]Game, 1)

	for scanner.Scan() {
		games = append(games, gameParser(scanner.Text()))
	}
	for i := range games {
		if games[i].Red <= MaxRed && games[i].Blue <= MaxBlue && games[i].Green <= MaxGreen {
			validIdCount += games[i].Id
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return validIdCount
}

func SecondPart(filePath string) int {
	setPower := 0

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	games := make([]Game, 1)

	for scanner.Scan() {
		games = append(games, gameParser(scanner.Text()))
	}
	for i := range games {
		setPower += games[i].Red * games[i].Blue * games[i].Green
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return setPower
}

func gameParser(gameString string) Game {
	var game *Game
	game = new(Game)
	var err error

	gameData := strings.Split(gameString, ": ")
	gameMetaData := gameData[0]
	gameSets := strings.Split(gameData[1], "; ")

	reNumb := regexp.MustCompile("(\\d+ )(green|blue|red)")

	game.Id, err = strconv.Atoi(strings.Split(gameMetaData, "Game ")[1])
	if err != nil {
		log.Fatal(err)
	}
	for i := range gameSets {

		array := reNumb.FindAllString(gameSets[i], -1)
		for i := range array {
			cubes := strings.Split(array[i], " ")
			count, err := strconv.Atoi(cubes[0])
			if err != nil {
				log.Fatal(err)
			}
			switch cubes[1] {
			case "green":
				if game.Green < count {
					game.Green = count
				}
			case "red":
				if game.Red < count {
					game.Red = count
				}
			case "blue":
				if game.Blue < count {
					game.Blue = count
				}
			}
		}
	}
	return *game
}
