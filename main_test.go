package main

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestFindNodeByID(t *testing.T) {
	var tree Tree
	tree.AddNode("1", "root", "")
	tree.AddNode("56", "5656", "1")
	tree.AddNode("2", "22", "1")
	tree.AddNode("3", "33", "2")
	tree.AddNode("4", "44", "2")

	var flagtests = []struct {
		id  string
		out bool
	}{
		{"", false},
		{"1", true},
		{"4", true},
		{"12333", false},
		{"2", true},
	}

	for i, tt := range flagtests {
		t.Run("case"+strconv.Itoa(i), func(t *testing.T) {
			node := findNodeByID(tree.root, tt.id)

			res := false
			if node != nil {
				res = true
			}
			if res != tt.out {
				t.Errorf("got %v, want %v\nTree%+v\n", res, tt.out, tree)
			}
		})
	}
}

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
		{"2", "root2", "", false},
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

func TestMoveNodeSuite(t *testing.T) {
	var tree Tree
	tree.AddNode("0", "root", "")
	tree.AddNode("1", "11", "0")
	tree.AddNode("4", "44", "0")
	tree.AddNode("5", "55", "4")
	tree.AddNode("6", "name", "4")

	tree.AddNode("2", "name", "1")
	tree.AddNode("3", "33", "2")
	tree.AddNode("33", "333", "1")

	var flagtests = []struct {
		id            string
		new_parent_id string
		out           bool
	}{
		{"", "", false},
		{"111", "", false},
		{"111", "1", false},
		{"1", "111", false},
		{"2", "4", false},
		{"1", "3", false},
		{"0", "5", false},
		{"1", "4", true},
		{"1", "0", true},
		{"1", "1", false},
	}

	for i, tt := range flagtests {
		t.Run("case"+strconv.Itoa(i), func(t *testing.T) {
			res := tree.MoveNode(tt.id, tt.new_parent_id)
			if res != tt.out {
				tree.PrintTree()
				t.Errorf("got %v, want %v\nTree%+v\n", res, tt.out, tree)
			}
		})
	}
}

func TestSliceFromTree(t *testing.T) {
	var tree Tree
	tree.AddNode("1", "root", "")
	tree.AddNode("2", "22", "1")
	tree.AddNode("3", "33", "1")

	n2 := Node{ID: "2", Name: "22", ParentID: "1", Children: nil}
	n3 := Node{ID: "3", Name: "33", ParentID: "1", Children: nil}
	n1 := Node{ID: "1", Name: "root", ParentID: "", Children: []*Node{&n2, &n3}}

	want := []*Node{&n1, &n2, &n3}
	got := sliceFromTree(tree.root)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("SliceFromTree: got %+v, want %+v\n", got, want)
	}
}

func TestSortNodes(t *testing.T) {
	nodes := []*Node{
		&Node{ID: "1", Name: "ccc"},
		&Node{ID: "2", Name: "bbb"},
		&Node{ID: "3", Name: "aaa"},
	}

	want := []*Node{
		&Node{ID: "3", Name: "aaa"},
		&Node{ID: "2", Name: "bbb"},
		&Node{ID: "1", Name: "ccc"},
	}

	sort.Sort(NameSorter(nodes))

	if !reflect.DeepEqual(nodes, want) {
		t.Errorf("SliceFromTree: got %+v, want %+v\n", nodes, want)
	}

}

func TestQueryMinMaxSuite(t *testing.T) {
	var tree Tree
	tree.AddNode("1", "root", "")
	tree.AddNode("2", "22", "1")
	tree.AddNode("3", "33", "1")
	tree.AddNode("4", "44", "1")
	tree.AddNode("5", "55", "2")
	tree.AddNode("6", "66", "2")
	tree.AddNode("9", "99", "6")

	tree.AddNode("7", "77", "4")
	tree.AddNode("8", "88", "7")

	var flagtests = []struct {
		mind  int
		maxd  int
		names []string
		ids   []string

		out []*Node
	}{
		{-1, -1, nil, nil, sliceFromTree(tree.root)},
		{-1, 100, nil, nil, sliceFromTree(tree.root)},
		{3, -1, nil, nil, []*Node{
			&Node{ID: "9", Name: "99", ParentID: "6"},
			&Node{ID: "8", Name: "88", ParentID: "7"}}},
		{100, -1, nil, nil, []*Node{}},
		{-1, 0, nil, nil, sliceFromTree(tree.root)[:1]},
		{0, 0, nil, nil, sliceFromTree(tree.root)[:1]},
		{3, 3, nil, nil, []*Node{
			&Node{ID: "9", Name: "99", ParentID: "6"},
			&Node{ID: "8", Name: "88", ParentID: "7"}}},
		{-1, -1, []string{"root"}, nil, sliceFromTree(tree.root)[:1]},
		{-1, -1, []string{"99", "88"}, nil, []*Node{
			&Node{ID: "9", Name: "99", ParentID: "6"},
			&Node{ID: "8", Name: "88", ParentID: "7"}}},
		{-1, -1, nil, []string{"9", "8"}, []*Node{
			&Node{ID: "9", Name: "99", ParentID: "6"},
			&Node{ID: "8", Name: "88", ParentID: "7"}}},
		{2, 2, nil, nil, []*Node{
			&Node{ID: "5", Name: "55", ParentID: "2"},
			&Node{ID: "6", Name: "66", ParentID: "2", Children: []*Node{&Node{ID: "9", Name: "99", ParentID: "6"}}},
			&Node{ID: "7", Name: "77", ParentID: "4", Children: []*Node{&Node{ID: "8", Name: "88", ParentID: "7"}}}}},
	}

	for i, tt := range flagtests {
		t.Run("case"+strconv.Itoa(i), func(t *testing.T) {
			res := tree.Query(tt.mind, tt.maxd, tt.names, tt.ids, nil)

			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("got %v len = %d, want %v len = %d\n", res, len(res), tt.out, len(tt.out))
				tree.PrintTree()
			}
		})
	}
}

func TestQueryNamesSuite(t *testing.T) {
	var tree Tree
	tree.AddNode("1", "root", "")
	tree.AddNode("2", "22", "1")
	tree.AddNode("3", "33", "1")
	tree.AddNode("4", "44", "1")

	var flagtests = []struct {
		mind  int
		maxd  int
		names []string
		ids   []string

		out []*Node
	}{
		{-1, -1, []string{"33"}, nil, []*Node{
			&Node{ID: "3", Name: "33", ParentID: "1"}}},

		{-1, -1, []string{"33", "44"}, nil, []*Node{
			&Node{ID: "3", Name: "33", ParentID: "1"},
			&Node{ID: "4", Name: "44", ParentID: "1"}}},
	}

	for i, tt := range flagtests {
		t.Run("case"+strconv.Itoa(i), func(t *testing.T) {
			res := tree.Query(tt.mind, tt.maxd, tt.names, tt.ids, nil)

			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("got %v len = %d, want %v len = %d\n", res, len(res), tt.out, len(tt.out))
				tree.PrintTree()
			}
		})
	}
}

func TestQuerySuite(t *testing.T) {
	var tree Tree
	tree.AddNode("1", "root", "")
	tree.AddNode("2", "22", "1")
	tree.AddNode("3", "33", "1")
	tree.AddNode("4", "44", "1")
	tree.AddNode("5", "55", "2")
	tree.AddNode("6", "66", "2")
	tree.AddNode("9", "99", "6")

	tree.AddNode("7", "77", "4")
	tree.AddNode("8", "88", "7")

	var flagtests = []struct {
		mind  int
		maxd  int
		names []string
		ids   []string

		out []*Node
	}{
		{-1, -1, []string{"55", "66", "77"}, nil, []*Node{
			&Node{ID: "5", Name: "55", ParentID: "2"},
			&Node{ID: "6", Name: "66", ParentID: "2", Children: []*Node{&Node{ID: "9", Name: "99", ParentID: "6"}}},
			&Node{ID: "7", Name: "77", ParentID: "4", Children: []*Node{&Node{ID: "8", Name: "88", ParentID: "7"}}}}},
		{2, 2, []string{"55"}, nil, []*Node{
			&Node{ID: "5", Name: "55", ParentID: "2"}}},
		{1, 3, []string{"root", "55", "99"}, nil, []*Node{
			&Node{ID: "5", Name: "55", ParentID: "2"},
			&Node{ID: "9", Name: "99", ParentID: "6"}}},
	}

	for i, tt := range flagtests {
		t.Run("case"+strconv.Itoa(i), func(t *testing.T) {
			res := tree.Query(tt.mind, tt.maxd, tt.names, tt.ids, nil)

			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("got %v len = %d, want %v len = %d\n", res, len(res), tt.out, len(tt.out))
				tree.PrintTree()
			}
		})
	}
}
