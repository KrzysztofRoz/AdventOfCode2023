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
	value     string
	leftNode  *Node
	RightNode *Node
}

type NodeData struct {
	value string
	left  string
	right string
}

type NodesDataArray []NodeData

func main() {
	path := "/home/krzysztof/Repos/AdventOfCode2023/inputs/input08.txt"
	fmt.Println(FirstPart(path))
	//fmt.Println(SecondPart(path))
}

func FirstPart(filePath string) int {
	result := 0
	emptyLineFlag := false
	var instruction []string
	nodesData := make([]NodeData, 0)
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if emptyLineFlag && len(scanner.Text()) != 0 {
			nodesData = append(nodesData, ParseNodeData(scanner.Text()))
		} else if len(scanner.Text()) != 0 {
			instruction = strings.Split(scanner.Text(), "")
		} else {
			emptyLineFlag = true
		}
	}
	fmt.Println(instruction)
	fmt.Println(nodesData)
	result = travel(instruction, nodesData)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func SecondPart(filePath string) int {
	result := 0

	return result
}

func ParseNodeData(nodeString string) NodeData {
	var nodeData *NodeData
	nodeData = new(NodeData)
	reg := regexp.MustCompile("[A-Z]{3}")

	data := reg.FindAllString(nodeString, -1)
	nodeData.value = data[0]
	nodeData.left = data[1]
	nodeData.right = data[2]

	return *nodeData

}

// func ParseNodes(nodesData NodesDataArray) Node {
// 	var root Node
// 	root.value = "AAA"
// 	root.leftNode = nodesData.GetNodeByValue("BBB")

// 	return root
// }
func (nodes NodesDataArray) GetNodeByValue(value string) NodeData {
	idx := -1
	for i := range nodes {
		if nodes[i].value == value {
			idx = i
			break
		}
	}
	return nodes[idx]
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
