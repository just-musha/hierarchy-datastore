package main

import (
	"reflect"
	"testing"
)

func TestAddRoot(t *testing.T) {
	var tree Tree
	result := tree.AddNode("1", "root", "")

	wanttree := Tree{root: &Node{
		ID:   "1",
		Name: "root",
	},
	}
	if result == false {
		t.Errorf("Cannot add Node")
	}
	if !reflect.DeepEqual(tree, wanttree) {
		t.Errorf("Add Root incorrect:\ngot %+v\nwant %+v\n", tree, wanttree)
	}
}

func TestAddNode(t *testing.T) {
	var tree Tree
	tree.AddNode("1", "root", "")
	tree.AddNode("2", "22", "1")
	tree.AddNode("56", "5656", "1")

	wanttree := Tree{root: &Node{
		ID:   "1",
		Name: "root",
		Children: []*Node{
			&Node{ID: "2", Name: "22", ParentID: "1"},
			&Node{ID: "56", Name: "5656", ParentID: "1"},
		},
	},
	}

	if !reflect.DeepEqual(tree, wanttree) {
		t.Errorf("Add Node incorrect:\ngot %+v\nwant %+v\n", tree.root, wanttree.root)
	}

}

func TestDeleteNode(t *testing.T) {
	var tree Tree
	tree.AddNode("1", "root", "")
	tree.AddNode("56", "5656", "1")
	tree.AddNode("2", "22", "1")

	tree.DeleteNode("56")

	wanttree := Tree{root: &Node{
		ID:   "1",
		Name: "root",
		Children: []*Node{
			&Node{ID: "2", Name: "22", ParentID: "1"},
		},
	},
	}

	if !reflect.DeepEqual(tree, wanttree) {
		t.Errorf("Wrong Delete Node:\ngot %+v\nwant %+v\n", tree, wanttree)
	}
}

func TestDeleteRoot(t *testing.T) {
	var tree Tree
	tree.AddNode("1", "root", "")

	tree.DeleteNode("1")

	wanttree := Tree{root: nil}

	if !reflect.DeepEqual(tree, wanttree) {
		t.Errorf("Wrong Delete Root:\ngot %+v\nwant %+v\n", tree, wanttree)
	}
}
