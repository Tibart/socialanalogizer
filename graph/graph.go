package graph

import (
	"encoding/json"
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
	fv.adjacent = append(fv.adjacent, tv)

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

// getVertex returns the vertex coresponding to the key if exists.
func (g *Graph) getVertex(key string) (*Vertex, error) {
	for _, v := range g.vertices {
		if v.key == key {
			return v, nil
		}
	}

	return nil, fmt.Errorf(errVertexExists, key)
}

func (g *Graph) String() string {
	j, _ := json.MarshalIndent(*g, "", "\t")
	return string(j)
}
