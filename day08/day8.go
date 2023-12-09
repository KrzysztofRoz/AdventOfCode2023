package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	value string
	left  string
	right string
}

type NodesDataArray []Node

func main() {
	path := "/home/krzysztof/Repos/AdventOfCode2023/inputs/input08.txt"
	//fmt.Println(FirstPart(path))
	fmt.Println(SecondPart(path))
}

func FirstPart(filePath string) int {
	result := 0
	emptyLineFlag := false
	var instruction []string
	nodes := make([]Node, 0)
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if emptyLineFlag && len(scanner.Text()) != 0 {
			nodes = append(nodes, ParseNode(scanner.Text()))
		} else if len(scanner.Text()) != 0 {
			instruction = strings.Split(scanner.Text(), "")
		} else {
			emptyLineFlag = true
		}
	}
	fmt.Println(instruction)
	fmt.Println(nodes)
	result = travel(instruction, nodes)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func SecondPart(filePath string) int {
	result := 0
	emptyLineFlag := false
	var instruction []string
	nodes := make([]Node, 0)

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if emptyLineFlag && len(scanner.Text()) != 0 {
			nodes = append(nodes, ParseNode(scanner.Text()))
		} else if len(scanner.Text()) != 0 {
			instruction = strings.Split(scanner.Text(), "")
		} else {
			emptyLineFlag = true
		}
	}
	fmt.Println(instruction)
	result = part2Travel(instruction, nodes)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func ParseNode(nodeString string) Node {
	var node *Node
	node = new(Node)
	reg := regexp.MustCompile("[A-Z]{3}")

	data := reg.FindAllString(nodeString, -1)
	node.value = data[0]
	node.left = data[1]
	node.right = data[2]

	return *node

}

func (nodes NodesDataArray) GetNodeByValue(value string) Node {
	idx := -1
	for i := range nodes {
		if nodes[i].value == value {
			idx = i
			break
		}
	}
	return nodes[idx]
}

func (nodes NodesDataArray) GetStartingNodes() []Node {
	startNodes := make([]Node, 0)
	reg := regexp.MustCompile("[A-Z]{2}A")
	for _, node := range nodes {
		if reg.Match([]byte(node.value)) {
			startNodes = append(startNodes, node)
		}
	}

	return startNodes
}

func travel(instruction []string, nodes NodesDataArray) int {
	result := 0
	findEnd := false
	insSize := len(instruction)
	var i int
	var direction string
	var right string
	var left string
	node := nodes.GetNodeByValue("AAA")
	right = node.right
	left = node.left

	for !findEnd {
		direction = instruction[i%insSize]
		if direction == "L" {
			node = nodes.GetNodeByValue(left)
		}
		if direction == "R" {
			node = nodes.GetNodeByValue(right)
		}
		right = node.right
		left = node.left
		result++
		i++
		if node.value == "ZZZ" {
			fmt.Println("Find ZZZ")
			findEnd = true
		}

	}

	return result
}
func part2Travel(instruction []string, nodes NodesDataArray) int {
	result := 0
	findEnd := false
	startingNodes := nodes.GetStartingNodes()
	insSize := len(instruction)
	stepToEnd := make([]int, len(startingNodes))
	reg := regexp.MustCompile("[A-Z]{2}Z")
	i := 0
	var direction string
	var right string
	var left string
	fmt.Println(startingNodes)

	for !findEnd {
		result++
		findEnd = true
		direction = instruction[i]

		for j := 0; j < len(startingNodes); j++ {
			right = startingNodes[j].right
			left = startingNodes[j].left
			if direction == "L" {
				startingNodes[j] = nodes.GetNodeByValue(left)
			}
			if direction == "R" {
				startingNodes[j] = nodes.GetNodeByValue(right)
			}
			if reg.MatchString(startingNodes[j].value) && stepToEnd[j] == 0 {
				stepToEnd[j] = result
			}
		}
		for k := range stepToEnd {
			fmt.Println(stepToEnd)
			if stepToEnd[k] != 0 {
				findEnd = true
			} else {
				findEnd = false
			}
		}
		i++
		i = i % insSize
	}
	result = LCM(stepToEnd)

	return result
}

func LCM(numbers []int) int {
	result := 0

	return result
}
