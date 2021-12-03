package graph

import (
	"fmt"
)

const (
	errVertexExists = "vertex with the key '%v' already exists"
)

// Graph is the top level struct for storing the graph data
type Graph struct {
	Vertices []*Vertex
}

// Vertex reassembles the object that has relations
type Vertex struct {
	Key      string
	Adjacent []*Vertex
}

// AddVertex add a new Vertex to the Graph
func (g *Graph) AddVertex(key string) error {
	// Check if Vertex already exists
	if g.Containes(key) {
		return fmt.Errorf(errVertexExists, key)
	}

	// Add Vertex to Graph
	g.Vertices = append(g.Vertices, &Vertex{key, nil})

	return nil
}

// AddEdge
func (g *Graph) AddEdge(from, to string) error {
	// TODO: Figure out if vertex can be made if nit exists
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

	// Add to vertex to from vertex
	fv.Adjacent = append(fv.Adjacent, tv)

	return nil
}

// exists checks if a vertex already exists in graph Vertices
func (g *Graph) Containes(key string) bool {
	for _, v := range g.Vertices {
		if v.Key == key {
			return true
		}
	}

	return false
}

// getVertex returns the vertex coresponding to the key if exists.
func (g *Graph) getVertex(key string) (*Vertex, error) {
	for _, v := range g.Vertices {
		if v.Key == key {
			return v, nil
		}
	}

	return nil, fmt.Errorf(errVertexExists, key)
}

func (g *Graph) Print() {
	for _, v := range g.Vertices {
		fmt.Printf("Vertex '%v' :", v.Key)
		for _, v := range v.Adjacent {
			fmt.Printf(" %v", v.Key)
		}
		fmt.Println()
	}
}
