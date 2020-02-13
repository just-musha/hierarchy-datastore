package main

import (
	hierarchy "hierarchy_wip/hierarchy"
)

func main() {
	var tree hierarchy.Tree

	tree.AddNode("1", "root", "")

	tree.AddNode("2", "22", "1")
	tree.AddNode("3", "33", "1")
	tree.AddNode("4", "44", "1")

	tree.AddNode("5", "55", "2")
	tree.AddNode("6", "66", "2")

	tree.AddNode("7", "77", "4")

	tree.PrintTree()

}
