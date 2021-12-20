package graph

import (
	"fmt"
)

const (
	errVertexExists = "vertex with the key '%v' already exists"
)

// Graph is the top level struct for storing the graph data
type Graph struct {
	vertices []*Vertex
}

// Vertex reassembles the object that has relations
type Vertex struct {
	key      string
	adjacent []*Vertex
}

func NewGraph() *Graph {
	return &Graph{}
}

// TODO implement export interface of struct's

// AddVertex add a new Vertex to the Graph
func (g *Graph) AddVertex(key string) error {
	// Check if Vertex already exists
	if g.Containes(key) {
		return fmt.Errorf(errVertexExists, key)
	}

	// Add Vertex to Graph
	g.vertices = append(g.vertices, &Vertex{key, nil})

	return nil
}

// AddEdge
func (g *Graph) AddEdge(from, to string, biDirectional bool) error {
	// Get from vertex
	fv, err := g.getVertex(from)
	if err != nil {
		return err
	}

	// Get to vertex
	tv, err := g.getVertex(to)
	if err != nil {
		return err
	}

	// Add edge to adjacents
	if !fv.Containes(tv.key) {
		fv.adjacent = append(fv.adjacent, tv)
	}
	if biDirectional && !tv.Containes(fv.key) {
		tv.adjacent = append(tv.adjacent, fv)
	}

	return nil
}

// exists checks if a vertex already exists in graph vertices
func (g *Graph) Containes(key string) bool {
	for _, v := range g.vertices {
		if v.key == key {
			return true
		}
	}
	return false
}

func (v *Vertex) Containes(key string) bool {
	for _, e := range v.adjacent {
		if e.key == key {
			return true
		}
	}
	return false
}

func (g *Graph) GetVerticesKeys() []string {
	k := []string{}
	for _, v := range g.vertices {
		k = append(k, v.key)
	}
	return k
}

// Export graph database as nodes and edges
func (g *Graph) Export() interface{} {
	type node struct {
		Id    string `json:"id"`
		Label string `json:"label"`
	}
	nd := []node{}
	type edge struct {
		From string `json:"from"`
		To   string `json:"to"`
	}
	ed := []edge{}

	// Add data to struct's
	for _, v := range g.vertices {
		nd = append(nd, node{v.key, v.key})
		for _, a := range v.adjacent {
			ed = append(ed, edge{v.key, a.key})
		}
	}

	//
	type data struct {
		Nodes []node `json:"nodes"`
		Edges []edge `json:"edges"`
	}
	e := data{nd, ed}

	return e
}

// getVertex returns the vertex coresponding to the key if exists.
func (g *Graph) getVertex(key string) (*Vertex, error) {
	for _, v := range g.vertices {
		if v.key == key {
			return v, nil
		}
	}

	return nil, fmt.Errorf(errVertexExists, key)
}
