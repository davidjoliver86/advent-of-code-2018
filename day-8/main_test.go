package main

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

const input = "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"

func makeScanner() *bufio.Scanner {
	reader := strings.NewReader(input)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	return scanner
}
func TestSpawn(t *testing.T) {
	node := Spawn(makeScanner())
	if len(node.Children) != 2 {
		t.Error("Node does not have two children")
	}
	expected := []int{1, 1, 2}
	if !(reflect.DeepEqual(expected, node.Metadata)) {
		t.Error("Metadata slice does not equal [1, 1, 2]")
	}
}

func TestSumMetadata(t *testing.T) {
	node := Spawn(makeScanner())
	if SumMetadata(node) != 138 {
		t.Error("Sum of metadata does not equal 138")
	}
}

func TestNodeValue(t *testing.T) {
	node := Spawn(makeScanner())
	nodeValue := NodeValue(node)
	if nodeValue != 66 {
		t.Errorf("Value of root node does not equal 66 (it equals %d)", nodeValue)
	}
}
