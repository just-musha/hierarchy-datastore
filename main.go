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

func findChildIdxByID(node *Node, id string) int {
	for i := range node.Children {
		if node.Children[i].ID == id {
			return i
		}
	}
	return -1
}

func removeChildByID(node *Node, id string) bool {
	idx := findChildIdxByID(node, id)
	if idx == -1 {
		return false
	}

	len := len(node.Children)
	copy(node.Children[idx:], node.Children[idx+1:])
	node.Children[len-1] = nil
	node.Children = node.Children[:len-1]
	return true
}

// Check if node "name" is already present among
// children of parent node
func isTakenChildName(parent *Node, childname string) bool {
	for _, ch := range parent.Children {
		if ch.Name == childname {
			return true
		}
	}
	return false
}

// TODO: "add_node"
func (tr *Tree) AddNode(id, name, parent_id string) bool {
	if name == "" || id == "" {
		fmt.Errorf("Name and ID must be specified and not empty strings")
		return false
	}
	if parent_id == "" && tr.root != nil {
		fmt.Errorf("There can only be one root node (i.e., a node without a parent)")
		return false
	} else if parent_id == "" && tr.root == nil {
		// Add root node
		tr.root = &Node{ID: id, Name: name}
		return true
	}
	if !existInSubtree(tr.root, parent_id) {
		fmt.Errorf("If specified, parent node must exist")
		return false
	}
	if existInSubtree(tr.root, id) {
		fmt.Errorf("No two nodes in the tree can have the same ID")
		return false
	}

	parent := findNodeByID(tr.root, parent_id)
	exist := isTakenChildName(parent, name)
	if exist {
		fmt.Errorf("Two sibling nodes cannot have the same name")
		return false
	}
	parent.Children = append(parent.Children, &Node{ID: id, Name: name, ParentID: parent_id})
	return true
}

// TODO: "delete_node"
func (tr *Tree) DeleteNode(id string) bool {
	if id == "" {
		fmt.Errorf("ID must be specified and not an empty string")
		return false
	}
	node := findNodeByID(tr.root, id)
	if node == nil {
		fmt.Errorf("Node does not exist")
		return false
	}

	if len(node.Children) != 0 {
		fmt.Errorf("Node must not have children")
		return false
	}

	pid := node.ParentID
	if pid == "" {
		// Delete Root
		tr.root = nil
		return true
	}

	pnode := findNodeByID(tr.root, pid)
	return removeChildByID(pnode, id)
}

// TODO: "move_node"
func (tr *Tree) MoveNode(id, new_parent_id string) bool {
	if id == "" || new_parent_id == "" {
		fmt.Errorf("MoveNode: ID and new parent ID must be specified and not empty strings")
		return false
	}
	node := findNodeByID(tr.root, id)
	if node == nil {
		fmt.Errorf("MoveNode: id Node not found")
		return false
	}
	newp := findNodeByID(tr.root, new_parent_id)
	if newp == nil {
		fmt.Errorf("MoveNode: new parent id Node not found")
		return false
	}
	if isTakenChildName(newp, node.Name) {
		fmt.Errorf("MoveNode: name of node already exist among new parent children")
		return false
	}
	// Check for cycle
	// cycle if new_parent is in subtree of node
	if findNodeByID(node, new_parent_id) != nil {
		fmt.Errorf("MoveNode: moving node should not create a cycle")
		return false
	}

	oldp := findNodeByID(tr.root, node.ParentID)
	if !removeChildByID(oldp, id) {
		return false
	}

	node.ParentID = new_parent_id
	newp.Children = append(newp.Children, node)
	return true
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

	fmt.Println("---------------------")
	tree.DeleteNode("8")
	tree.DeleteNode("7")
	tree.PrintTree()
}
