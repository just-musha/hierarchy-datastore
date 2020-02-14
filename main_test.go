package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	hierarchy "hierarchy-datastore/hierarchy"
	"testing"
)

var node1 = hierarchy.Node{
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

var node2 = hierarchy.Node{
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

func TestResponse(t *testing.T) {
	var rq ResponseQuery
	rq.Nodes = []*hierarchy.Node{&node1, &node2}
	got, _ := json.Marshal(rq)
	want := "{\"nodes\":[{\"name\":\"11111\",\"id\":\"1\",\"parent_id\":\"5\"},{\"name\":\"222\",\"id\":\"2\",\"parent_id\":\"6\"}]}"
	if string(got) != want {
		t.Errorf("got %v, want %v\n", string(got), want)
	}
}

func TestResponseQuery(t *testing.T) {
	var resp Response
	got, _ := json.Marshal(resp)
	want := "{\"ok\":false}"
	if string(got) != want {
		t.Errorf("got %v, want %v\n", string(got), want)
	}
}

func TestRequestAdd(t *testing.T) {
	var got Request
	got.Query.SetDefaults()

	jsonReq := "{\"add_node\":{\"id\":\"1\",\"name\":\"Root\"}}"

	err := json.Unmarshal([]byte(jsonReq), &got)
	if err != nil {
		t.Errorf("Cannot Unmarshal: %v\n", err)
	}

	var want Request
	want.Query.SetDefaults()
	want.AddNode = AddNodeParams{ID: "1", Name: "Root", ParentID: ""}

	if got.AddNode != want.AddNode {
		t.Errorf("\ngot %+v\nwant %+v\n", got, want)
	}
}

func TestRequestQuery(t *testing.T) {
	var got Request
	got.Query.SetDefaults()

	jsonReq := "{\"query\":{\"min_depth\":2,\"names\":[\"B\"]}}"

	err := json.Unmarshal([]byte(jsonReq), &got)
	if err != nil {
		t.Errorf("Cannot Unmarshal: %v\n", err)
	}

	var want Request
	want.Query.SetDefaults()
	want.Query.MinD = 2
	want.Query.Names = []string{"B"}

	fmt.Printf("got:%+v\nwant:%+v\n", got, want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v\n", got, want)
	}
}
