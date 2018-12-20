package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	util "github.com/davidjoliver86/advent-of-code-2018"
)

type Node struct {
	Children []*Node
	Metadata []int
}

func Spawn(scanner *bufio.Scanner) *Node {
	n := &Node{}

	// pick two off scanner - 1) how many children, 2) how many metadata points
	scanner.Scan()
	children, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	metadata, _ := strconv.Atoi(scanner.Text())
	n.Children = make([]*Node, children)
	n.Metadata = make([]int, metadata)
	for i := 0; i < len(n.Children); i++ {
		n.Children[i] = Spawn(scanner)
	}
	for i := 0; i < len(n.Metadata); i++ {
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		n.Metadata[i] = m
	}
	return n
}

func SumMetadata(n *Node) int {
	i := 0
	for _, m := range n.Metadata {
		i += m
	}
	for _, child := range n.Children {
		i += SumMetadata(child)
	}
	return i
}

func NodeValue(n *Node) int {
	// if no child nodes, just return sum of metadata
	if len(n.Children) == 0 {
		return util.Sum(n.Metadata)
	}

	// otherwise each metadata entry is a 1-based index of this node's children
	// sum up the values of those children
	// if an index does not exist, pass
	value := 0
	length := len(n.Children)
	for _, index := range n.Metadata {
		if index <= length {
			value += NodeValue(n.Children[index-1])
		}
	}
	return value
}

func main() {
	fp, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanWords)
	root := Spawn(scanner)
	fmt.Println(SumMetadata(root))
	fmt.Println(NodeValue(root))
}
