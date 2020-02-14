package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
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
	MinD    *int     `json:"min_depth"`
	MaxD    *int     `json:"max_depth"`
	Names   []string `json:"names"`
	IDs     []string `json:"ids"`
	RootIDs []string `json:"root_ids"`
}

type Request struct {
	AddNode    *AddNodeParams    `json:"add_node"`
	DeleteNode *DeleteNodeParams `json:"delete_node"`
	MoveNode   *MoveNodeParams   `json:"move_node"`
	Query      *QueryParams      `json:"query"`
}

type Response struct {
	Ok bool `json:"ok"`
}
type ResponseQuery struct {
	Nodes []*hierarchy.Node `json:"nodes"`
}

func writeResponse(status bool) {
	resp, err := json.Marshal(Response{Ok: status})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))
}

func writeResponseNodes(nodes []*hierarchy.Node) {
	resp, err := json.Marshal(ResponseQuery{Nodes: nodes})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))
}

func AnalyseReqest(tree *hierarchy.Tree, req string) {
	var request Request

	err := json.Unmarshal([]byte(req), &request)
	if err != nil {
		panic(err)
	}
	if request.AddNode != nil {
		r := *request.AddNode
		fmt.Fprintf(os.Stderr, ">> Request Add: %q\n", request.AddNode)
		status := tree.AddNode(r.ID, r.Name, r.ParentID)
		writeResponse(status)

	} else if request.DeleteNode != nil {
		r := *request.DeleteNode
		fmt.Fprintf(os.Stderr, ">> Request Delete: %q\n", request.DeleteNode)
		status := tree.DeleteNode(r.ID)
		writeResponse(status)

	} else if request.MoveNode != nil {
		r := *request.MoveNode
		fmt.Fprintf(os.Stderr, ">> Request Move: %q\n", request.MoveNode)
		status := tree.DeleteNode(r.ID)
		writeResponse(status)

	} else if request.Query != nil {
		r := *request.Query
		fmt.Fprintf(os.Stderr, ">> Request Query: %+v\n", request.Query)
		mind, maxd := -1, -1
		if r.MinD != nil {
			mind = *r.MinD
		}
		if r.MaxD != nil {
			maxd = *r.MaxD
		}
		result := tree.Query(mind, maxd, r.Names, r.IDs, r.RootIDs)
		writeResponseNodes(result)
	} else {
		panic("unknown request")
	}
}

func main() {

	var tree hierarchy.Tree

	//scanner := bufio.NewScanner(strings.NewReader("{\"add_node\":{\"id\":\"1\",\"name\":\"Root\"}}\n{\"query\":{\"min_depth\":2,\"names\":[\"B\"]}}"))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		AnalyseReqest(&tree, scanner.Text())
	}
}
