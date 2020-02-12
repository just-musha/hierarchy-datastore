package main

import (
	"reflect"
	"strconv"
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

func TestAddNodeSuite(t *testing.T) {
	var flagtests = []struct {
		id        string
		name      string
		parent_id string
		out       bool
	}{
		{"", "", "", false},
		{"1", "root", "", true},
		{"", "1", "", false},
		{"", "", "1", false},
		{"1", "11", "111", false},
	}

	var tree Tree

	for i, tt := range flagtests {
		t.Run("case"+strconv.Itoa(i), func(t *testing.T) {
			res := tree.AddNode(tt.id, tt.name, tt.parent_id)
			if res != tt.out {
				t.Errorf("got %v, want %v\nTree%+v\n", res, tt.out, tree)

				if tt.out == true {
					tree.DeleteNode(tt.id)
				}
			}
		})
	}
}

func TestDeleteNodeSuite(t *testing.T) {
	var flagtests = []struct {
		id  string
		out bool
	}{
		{"", false},
		{"1", false},
		{"123", false},
		{"0", false},
		{"2", true},
		{"1", true},
		{"0", true},
	}

	var tree Tree
	tree.AddNode("0", "root", "")
	tree.AddNode("1", "11", "0")
	tree.AddNode("2", "22", "1")

	for i, tt := range flagtests {
		t.Run("case"+strconv.Itoa(i), func(t *testing.T) {
			res := tree.DeleteNode(tt.id)
			if res != tt.out {
				t.Errorf("got %v, want %v\nTree%+v\n", res, tt.out, tree)
			}
		})
	}

}
