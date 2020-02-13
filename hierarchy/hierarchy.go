package hierarchy

import (
	"fmt"
	"sort"
)

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

type NameSorter []*Node

func (nodes NameSorter) Len() int           { return len(nodes) }
func (nodes NameSorter) Swap(i, j int)      { nodes[i], nodes[j] = nodes[j], nodes[i] }
func (nodes NameSorter) Less(i, j int) bool { return nodes[i].Name < nodes[j].Name }

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

	// Keep children sorted by node.Name
	sort.Sort(NameSorter(parent.Children))

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

	// Keep children sorted by node.Name
	sort.Sort(NameSorter(newp.Children))
	return true
}

// Check if node satisfy all criteria
func (node *Node) check(result *[]*Node, depth int, min_depth, max_depth int, names, ids []string) {
	ok := true
	if min_depth != -1 && depth < min_depth {
		ok = false
	}

	if max_depth != -1 && depth > max_depth {
		ok = false
	}

	if len(names) != 0 {
		inlist := false
		for _, n := range names {
			if node.Name == n {
				inlist = true
			}
		}
		if !inlist {
			ok = false
		}
	}

	if len(ids) != 0 {
		inlist := false
		for _, id := range ids {
			if node.ID == id {
				inlist = true
			}
		}
		if !inlist {
			ok = false
		}
	}

	if ok {
		//fmt.Printf("Appending Node %v\n", node)
		*result = append(*result, node)
	}

	for _, ch := range node.Children {
		ch.check(result, depth+1, min_depth, max_depth, names, ids)
	}
}

func sliceFromTree_internal(root *Node, result *[]*Node) {
	*result = append(*result, root)
	for i := range root.Children {
		sliceFromTree_internal(root.Children[i], result)
	}
}

// Make slice from subtree
func sliceFromTree(root *Node) []*Node {
	var result []*Node
	sliceFromTree_internal(root, &result)
	return result
}

// Helper for Query function, removes root IDs which belong to any other subtree
func (tr Tree) filterRootIDs(root_ids []string) []string {
	result := []string{}
	for i := range root_ids {

		checknode := findNodeByID(tr.root, root_ids[i])
		inAnySubtree := false

		for j := range root_ids {
			if i == j {
				continue
			}
			rootnode := findNodeByID(tr.root, root_ids[j])
			if rootnode != nil && checknode != nil && existInSubtree(rootnode, checknode.ID) {
				inAnySubtree = true
			}
		}

		if !inAnySubtree {
			result = append(result, checknode.ID)
		}
	}
	return result
}

func (tr Tree) Query(min_depth, max_depth int, names, ids, root_ids []string) []*Node {

	result := []*Node{}

	if len(root_ids) == 0 {
		tr.root.check(&result, 0, min_depth, max_depth, names, ids)
	} else {

		roots := tr.filterRootIDs(root_ids)

		for _, id := range roots {
			node := findNodeByID(tr.root, id)
			if node != nil {
				node.check(&result, 0, min_depth, max_depth, names, ids)
			}
		}
	}

	return result
}
