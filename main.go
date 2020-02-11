package main

import "fmt"

type Node struct {
	Name     string
	ID       string
	ParentID string
	Children []*Node
}

type Tree struct {
	root *Node
}

func (tr Tree) PrintTree() {
	tr.root.print("")
}

func (node Node) print(prefix string) {
	fmt.Printf(prefix+">>id = %s, name = %s, parent_id = %s, child num = %d\n",
		node.ID, node.Name, node.ParentID, len(node.Children))
	for _, ch := range node.Children {
		ch.print(prefix + "\t")
	}
}

func existInSubtree(root *Node, id string) bool {
	return findNodeByID(root, id) != nil
}

func traverse(root *Node, found **Node, id string) {
	if root.ID == id {
		*found = root
		return
	}
	for _, ch := range root.Children {
		traverse(ch, found, id)
	}
}

func findNodeByID(root *Node, id string) *Node {
	var result *Node
	traverse(root, &result, id)
	return result

}

// TODO: "add_node"
func (tr Tree) AddNode(id, name, parent_id string) bool {
	if parent_id == "" && tr.root != nil {
		fmt.Errorf("There can only be one root node (i.e., a node without a parent)")
		return false
	}
	if !existInSubtree(tr.root, parent_id) {
		fmt.Errorf("If specified, parent node must exist")
		return false
	}
	if name == "" || id == "" {
		fmt.Errorf("Name and ID must be specified and not empty strings")
		return false
	}
	if existInSubtree(tr.root, id) {
		fmt.Errorf("No two nodes in the tree can have the same ID")
		return false
	}
	parent := findNodeByID(tr.root, parent_id)
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
	tree := Tree{root: &Node{ID: "1", Name: "root"}}

	tree.AddNode("2", "22", "1")
	tree.AddNode("3", "33", "1")
	tree.AddNode("4", "44", "1")

	tree.AddNode("5", "55", "2")
	tree.AddNode("6", "66", "2")

	tree.AddNode("7", "77", "4")

	node := findNodeByID(tree.root, "5")
	fmt.Printf("Found Node = %+v\n", node)

	tree.AddNode("8", "88", "7")

	tree.PrintTree()
}
