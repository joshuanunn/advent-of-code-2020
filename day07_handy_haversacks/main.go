package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	outer = regexp.MustCompile(`(\w+\s\w+) bags contain`)
	inner = regexp.MustCompile(`([0-9]) (\w+\s\w+) bag[s]?`)
)

type Node struct {
	desc     string
	qty      int
	children []*Node
}

type Tree struct {
	root *Node
}

func main() {
	// Read inputs into map of []*Node representing bag rules
	rules := readInputs("input.txt")

	// Part 1 - count number of top level bags containing a shiny gold bag somewhere inside
	// construct a slice of *Trees to represent bags
	var bags []*Tree
	for desc := range rules {
		bags = append(bags, createTree(rules, desc))
	}
	count := searchTrees(bags, "shiny gold")
	fmt.Printf("Part 1 - total count of top level bags containing a shiny gold bag within is %d.\n", count)

	// Part 2 - count total number of bags nested inside a shiny gold bag
	// construct a shiny gold bag
	shinyGold := createTree(rules, "shiny gold")
	count = total(shinyGold.root) - 1
	fmt.Printf("Part 2 - total count of bags within a shiny gold bag is %d.\n", count)
}

func (n *Node) insert(node *Node) {
	if n == nil {
		return
	}
	n.children = append(n.children, &Node{desc: node.desc, qty: node.qty})
}

func fill(n *Node, rules map[string][]*Node) {
	bags := rules[n.desc]
	if bags == nil {
		return
	}
	for _, bag := range bags {
		n.insert(bag)
	}
	for _, c := range n.children {
		fill(c, rules)
	}
}

func search(n *Node, target string) int {
	var count int
	if n.desc == target {
		return count + 1
	}
	if n.children == nil {
		return count
	}
	for _, c := range n.children {
		count += search(c, target)
	}
	return count
}

func total(n *Node) int {
	if n.children == nil {
		return n.qty
	}
	var count int
	for _, c := range n.children {
		count += total(c)
	}
	return count*n.qty + n.qty
}

func createTree(rules map[string][]*Node, key string) *Tree {
	// Setup top level tree
	tree := &Tree{}
	tree.root = &Node{desc: key, qty: 1}

	// Recursively fill tree with bags using rules
	fill(tree.root, rules)
	return tree
}

func searchTrees(trees []*Tree, target string) int {
	// Recursively search tree for target bag
	var count int
	for _, tree := range trees {
		// Ignore top level containing bag
		if tree.root.desc == target {
			continue
		}
		if search(tree.root, target) > 0 {
			count++
		}
	}
	return count
}

// Print all nodes in a tree, indenting each sublevel
func printTree(n *Node, indent int) {
	if n == nil {
		return
	}
	for i := 0; i < indent; i++ {
		fmt.Print("-")
	}
	fmt.Printf("%s (%d)\n", n.desc, n.qty)
	for _, n := range n.children {
		printTree(n, indent+1)
	}
}

// Read in file and parse into a map of bag rules
func readInputs(filename string) map[string][]*Node {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to open input.txt")
	}
	lines := string(b)

	var rules = make(map[string][]*Node)
	for _, line := range strings.Split(lines, "\n") {
		bags := []*Node{}
		desc := outer.FindStringSubmatch(line)[1]
		matches := inner.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			qty, err := strconv.Atoi(match[1])
			if err != nil {
				log.Fatalf("failed to parse int")
			}
			bags = append(bags, &Node{desc: match[2], qty: qty})
		}
		rules[desc] = bags
	}
	return rules
}
