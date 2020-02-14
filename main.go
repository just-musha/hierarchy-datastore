package main

import (
	"encoding/json"
	"fmt"
	hierarchy "hierarchy-datastore/hierarchy"
)

type AddNodeParams struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}

type DeleteNodeParams struct {
	ID string `json:"id"`
}

type MoveNodeParams struct {
	ID          string `json:"id"`
	NewParentID string `json:"new_parent_id"`
}

type QueryParams struct {
	MinD    int      `json:"min_depth"`
	MaxD    int      `json:"max_depth"`
	Names   []string `json:"names"`
	IDs     []string `json:"ids"`
	RootIDs []string `json:"root_ids"`
}

func (qp *QueryParams) SetDefaults() {
	qp.MinD = -1
	qp.MaxD = -1
}

type Request struct {
	AddNode    AddNodeParams    `json:"add_node"`
	DeleteNode DeleteNodeParams `json:"delete_node"`
	MoveNode   MoveNodeParams   `json:"move_node"`
	Query      QueryParams      `json:"query"`
}

type Response struct {
	Ok bool `json:"ok"`
}
type ResponseQuery struct {
	Nodes []*hierarchy.Node `json:"nodes"`
}

func main() {
	//var tree hierarchy.Tree

	node1 := hierarchy.Node{
		ID:       "1",
		Name:     "11111",
		ParentID: "5",
		Children: []*hierarchy.Node{
			&hierarchy.Node{
				Name:     "777",
				ID:       "7",
				ParentID: "8",
			},
		},
	}
	node2 := hierarchy.Node{
		ID:       "2",
		Name:     "222",
		ParentID: "6",
		Children: []*hierarchy.Node{
			&hierarchy.Node{
				Name:     "333",
				ID:       "4",
				ParentID: "9",
			},
		},
	}

	result, err := json.Marshal(node1)
	if err != nil {
		fmt.Printf("Cannot Marshal: %v\n", err)
	}
	fmt.Printf("result =%v\n", string(result))

	var rq ResponseQuery
	rq.Nodes = []*hierarchy.Node{&node1, &node2}
	resarr, err := json.Marshal(rq)
	if err != nil {
		fmt.Printf("Cannot Marshal: %v\n", err)
	}
	fmt.Println("nodes = ", string(resarr))

	var resp Response
	result, err = json.Marshal(resp)
	fmt.Println("array = ", string(result))

}
