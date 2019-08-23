package workflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type NodeTest struct {
	Key string
}

func (n NodeTest) ID() string { return n.Key }

func defaultGraph() *Graph {
	return &Graph{
		Nodes: []Node{
			NodeTest{Key: "nodeKey1"},
			NodeTest{Key: "nodeKey2"},
			NodeTest{Key: "nodeKey3"},
			NodeTest{Key: "nodeKey4"},
			NodeTest{Key: "nodeKey5"},
			NodeTest{Key: "nodeKey6"},
			NodeTest{Key: "nodeKey7"}},
		Edges: []Edge{
			{Src: "nodeKey1", Dst: "nodeKey2"},
			{Src: "nodeKey2", Dst: "nodeKey3"},
			{Src: "nodeKey2", Dst: "nodeKey4"},
			{Src: "nodeKey3", Dst: "nodeKey5"},
			{Src: "nodeKey4", Dst: "nodeKey6"},
			{Src: "nodeKey4", Dst: "nodeKey7"},
		},
	}
}

func TestChildrenIDs(t *testing.T) {
	var tests = []struct {
		graph    *Graph
		node     string
		children []string
	}{
		{graph: defaultGraph(), node: "nodeKey1", children: []string{"nodeKey2"}},
		{graph: defaultGraph(), node: "nodeKey2", children: []string{"nodeKey3", "nodeKey4"}},
		{graph: defaultGraph(), node: "nodeKey3", children: []string{"nodeKey5"}},
		{graph: defaultGraph(), node: "nodeKey4", children: []string{"nodeKey6", "nodeKey7"}},
		{graph: defaultGraph(), node: "nodeKey5", children: []string{}},
		{graph: defaultGraph(), node: "nodeKey6", children: []string{}},
		{graph: defaultGraph(), node: "nodeKey7", children: []string{}},
	}
	for _, test := range tests {
		assert.Equal(t, test.graph.ChildrenIDs(test.node), test.children)
	}
}

func TestParentIDs(t *testing.T) {
	var tests = []struct {
		graph   *Graph
		node    string
		parents []string
	}{
		{graph: defaultGraph(), node: "nodeKey1", parents: []string{}},
		{graph: defaultGraph(), node: "nodeKey2", parents: []string{"nodeKey1"}},
		{graph: defaultGraph(), node: "nodeKey3", parents: []string{"nodeKey2"}},
		{graph: defaultGraph(), node: "nodeKey4", parents: []string{"nodeKey2"}},
		{graph: defaultGraph(), node: "nodeKey5", parents: []string{"nodeKey3"}},
		{graph: defaultGraph(), node: "nodeKey6", parents: []string{"nodeKey4"}},
		{graph: defaultGraph(), node: "nodeKey7", parents: []string{"nodeKey4"}},
	}
	for _, test := range tests {
		assert.Equal(t, test.graph.ParentIDs(test.node), test.parents)
	}
}

func TestFindNode(t *testing.T) {
	var tests = []struct {
		graph   *Graph
		node    string
		present bool
	}{
		{graph: defaultGraph(), node: "nodeKey1", present: true},
		{graph: defaultGraph(), node: "nodeKey2", present: true},
		{graph: defaultGraph(), node: "nodeKey3", present: true},
		{graph: defaultGraph(), node: "nodeKey4", present: true},
		{graph: defaultGraph(), node: "nodeKey5", present: true},
		{graph: defaultGraph(), node: "nodeKey6", present: true},
		{graph: defaultGraph(), node: "nodeKey7", present: true},
		{graph: defaultGraph(), node: "nodeKey8", present: false},
	}
	for _, test := range tests {
		node, err := test.graph.FindNode(test.node)
		if test.present {
			assert.NoError(t, err)
			assert.Equal(t, node.ID(), test.node)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestHasNodes(t *testing.T) {
	var tests = []struct {
		graph    *Graph
		hasNodes bool
	}{
		{graph: defaultGraph(), hasNodes: true},
		{graph: &Graph{}, hasNodes: false},
	}
	for _, test := range tests {
		assert.Equal(t, test.graph.hasNodes(), test.hasNodes)
	}
}

func TestIsAcyclic(t *testing.T) {
	var tests = []struct {
		graph   *Graph
		acyclic bool
	}{
		{graph: defaultGraph(), acyclic: true},
		{graph: &Graph{
			Nodes: []Node{
				NodeTest{Key: "nodeKey1"},
				NodeTest{Key: "nodeKey2"},
			},
			Edges: []Edge{
				{Src: "nodeKey1", Dst: "nodeKey2"},
				{Src: "nodeKey2", Dst: "nodeKey1"},
			},
		}, acyclic: false},
		{graph: &Graph{
			Nodes: []Node{
				NodeTest{Key: "nodeKey1"},
				NodeTest{Key: "nodeKey2"},
				NodeTest{Key: "nodeKey3"},
			},
			Edges: []Edge{
				{Src: "nodeKey1", Dst: "nodeKey2"},
				{Src: "nodeKey2", Dst: "nodeKey3"},
				{Src: "nodeKey3", Dst: "nodeKey1"},
			},
		}, acyclic: false},
		{graph: &Graph{
			Nodes: []Node{
				NodeTest{Key: "nodeKey1"},
				NodeTest{Key: "nodeKey2"},
				NodeTest{Key: "nodeKey3"},
			},
			Edges: []Edge{
				{Src: "nodeKey1", Dst: "nodeKey2"},
				{Src: "nodeKey2", Dst: "nodeKey3"},
				{Src: "nodeKey3", Dst: "nodeKey2"},
			},
		}, acyclic: false},
		{graph: &Graph{
			Nodes: []Node{
				NodeTest{Key: "nodeKey1"},
				NodeTest{Key: "nodeKey2"},
				NodeTest{Key: "nodeKey3"},
				NodeTest{Key: "nodeKey4"},
			},
			Edges: []Edge{
				{Src: "nodeKey1", Dst: "nodeKey2"},
				{Src: "nodeKey1", Dst: "nodeKey3"},
				{Src: "nodeKey2", Dst: "nodeKey4"},
				{Src: "nodeKey3", Dst: "nodeKey4"},
			},
		}, acyclic: true},
	}
	for _, test := range tests {
		assert.Equal(t, test.graph.isAcyclic(), test.acyclic)
	}
}

func TestIsConnected(t *testing.T) {
	var tests = []struct {
		graph     *Graph
		node      string
		connected bool
	}{
		{graph: defaultGraph(), connected: true},
		{graph: &Graph{
			Nodes: []Node{
				NodeTest{Key: "nodeKey1"},
				NodeTest{Key: "nodeKey2"},
				NodeTest{Key: "nodeKey3"},
				NodeTest{Key: "nodeKey4"},
			},
			Edges: []Edge{
				{Src: "nodeKey1", Dst: "nodeKey2"},
				{Src: "nodeKey3", Dst: "nodeKey4"},
			},
		}, connected: false},
	}
	for i, test := range tests {
		assert.Equal(t, test.graph.isConnected(), test.connected, i)
	}
}

func TestVisitChildren(t *testing.T) {
	var tests = []struct {
		graph    *Graph
		node     string
		children []string
	}{
		{graph: defaultGraph(), node: "nodeKey1", children: []string{"nodeKey2", "nodeKey3", "nodeKey4", "nodeKey5", "nodeKey6", "nodeKey7"}},
		{graph: defaultGraph(), node: "nodeKey2", children: []string{"nodeKey3", "nodeKey4", "nodeKey5", "nodeKey6", "nodeKey7"}},
		{graph: defaultGraph(), node: "nodeKey3", children: []string{"nodeKey5"}},
		{graph: defaultGraph(), node: "nodeKey4", children: []string{"nodeKey6", "nodeKey7"}},
		{graph: defaultGraph(), node: "nodeKey5", children: []string{}},
		{graph: defaultGraph(), node: "nodeKey6", children: []string{}},
		{graph: defaultGraph(), node: "nodeKe7", children: []string{}},
	}
	for _, test := range tests {
		visit := make(map[string]bool)
		test.graph.dfs(test.node, func(node string) {
			visit[node] = true
		})
		for _, child := range test.children {
			assert.True(t, visit[child])
		}
	}
}

func TestGetRoot(t *testing.T) {
	var tests = []struct {
		graph *Graph
		node  string
		root  string
	}{
		{graph: defaultGraph(), node: "nodeKey1", root: "nodeKey1"},
		{graph: defaultGraph(), node: "nodeKey5", root: "nodeKey1"},
		{graph: defaultGraph(), node: "nodeKey6", root: "nodeKey1"},
		{graph: defaultGraph(), node: "nodeKey4", root: "nodeKey1"},
	}
	for _, test := range tests {
		assert.Equal(t, test.graph.getRoot(test.node), test.root)
	}
}

func TestIsMonoParental(t *testing.T) {
	var tests = []struct {
		graph *Graph
		max   int
	}{
		{graph: defaultGraph(), max: 1},
		{graph: &Graph{
			Nodes: []Node{
				NodeTest{Key: "nodeKey1"},
				NodeTest{Key: "nodeKey2"},
				NodeTest{Key: "nodeKey3"},
			},
			Edges: []Edge{
				{Src: "nodeKey1", Dst: "nodeKey3"},
				{Src: "nodeKey2", Dst: "nodeKey3"},
			},
		}, max: 2},
		{graph: &Graph{
			Nodes: []Node{
				NodeTest{Key: "nodeKey1"},
				NodeTest{Key: "nodeKey2"},
				NodeTest{Key: "nodeKey3"},
				NodeTest{Key: "nodeKey4"},
			},
			Edges: []Edge{
				{Src: "nodeKey1", Dst: "nodeKey2"},
				{Src: "nodeKey1", Dst: "nodeKey3"},
				{Src: "nodeKey2", Dst: "nodeKey4"},
				{Src: "nodeKey3", Dst: "nodeKey4"},
			},
		}, max: 2},
		{graph: &Graph{
			Nodes: []Node{
				NodeTest{Key: "nodeKey1"},
				NodeTest{Key: "nodeKey2"},
				NodeTest{Key: "nodeKey3"},
				NodeTest{Key: "nodeKey4"},
			},
			Edges: []Edge{
				{Src: "nodeKey1", Dst: "nodeKey4"},
				{Src: "nodeKey2", Dst: "nodeKey4"},
				{Src: "nodeKey3", Dst: "nodeKey4"},
			},
		}, max: 3},
	}
	for _, test := range tests {
		assert.Equal(t, test.max, test.graph.maximumParents())
	}
}
