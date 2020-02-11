package main

import "fmt"

type Node struct {
	Name     string
	ID       string
	ParentID string
	Children []*Node
}

var root *Node

func printTree(root *Node, prefix string) {
	fmt.Printf(prefix+">>id = %s, name = %s, parent_id = %s, child num = %d\n",
		root.ID, root.Name, root.ParentID, len(root.Children))
	for _, ch := range root.Children {
		printTree(ch, prefix+"\t")
	}
}

func existInSubtree(root *Node, id string) bool {
	return findNodeByID(root, id) != nil
}

func traverse(root, found *Node, id string) {
	if root.ID == id {
		*found = *root
		return
	}
	for _, ch := range root.Children {
		traverse(ch, found, id)
	}
}

func findNodeByID(root *Node, id string) *Node {
	var result Node
	traverse(root, &result, id)
	if result.ID == "" {
		return nil
	} else {
		return &result
	}
}

// TODO: "add_node"
func AddNode(id, name, parent_id string) bool {
	if parent_id == "" && root != nil {
		fmt.Errorf("There can only be one root node (i.e., a node without a parent)")
		return false
	}
	if !existInSubtree(root, parent_id) {
		fmt.Errorf("If specified, parent node must exist")
		return false
	}
	if name == "" || id == "" {
		fmt.Errorf("Name and ID must be specified and not empty strings")
		return false
	}
	if existInSubtree(root, id) {
		fmt.Errorf("No two nodes in the tree can have the same ID")
		return false
	}
	parent := findNodeByID(root, parent_id)
	exist := false
	for _, ch := range parent.Children {
		if ch.Name == name {
			exist = true
		}
	}
	if exist {
		fmt.Errorf("Two sibling nodes cannot have the same name")
		return false
	}
	parent.Children = append(parent.Children, &Node{ID: id, Name: name, ParentID: parent_id})
	fmt.Printf("Parent = %+q\n\n", parent)
	return true
}

// TODO: "delete_node"
func DeleteNode(id string) bool {
	return false
}

// TODO: "move_node"
func MoveNode(id, new_parent_id string) bool {
	return false
}

// TODO: "query"
func Query(min_depth, max_depth int, names, ids, root_ids []string) {
	return
}

func main() {
	root = &Node{ID: "1", Name: "root"}

	root.Children = append(root.Children, &Node{ID: "2", Name: "22", ParentID: "1"})
	root.Children = append(root.Children, &Node{ID: "3", Name: "33", ParentID: "1"})

	root.Children = append(root.Children, &Node{ID: "4", Name: "44", ParentID: "1"})

	ch0 := root.Children[0]
	ch2 := root.Children[2]

	ch0.Children = append(ch0.Children, &Node{ID: "5", Name: "55", ParentID: "2"})
	ch0.Children = append(ch0.Children, &Node{ID: "6", Name: "66", ParentID: "2"})

	ch2.Children = append(ch2.Children, &Node{ID: "7", Name: "77", ParentID: "4"})

	res := existInSubtree(root, "55")
	fmt.Printf("Found Node = %+v\n", res)

	res = AddNode("8", "88", "7")
	fmt.Println("Add status = ", res)

	printTree(root, "")
}
