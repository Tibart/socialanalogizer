package graph

import (
	"fmt"
)

const (
	errVertexNotExists = "vertex with the key '%v' already exists"
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
func (g *Graph) AddVertex(k string) error {
	// Check if Vertex already exists
	if g.exists(k) {
		return fmt.Errorf(errVertexNotExists, k)
	}

	// Add Vertex to Graph
	g.vertices = append(g.vertices, &Vertex{k, nil})

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
	tv, err := g.getVertex(from)
	if err != nil {
		return err
	}

	// Add to vertex to from vertex
	fv.adjacent = append(fv.adjacent, tv)

	return nil
}

// getVertex returns the vertex coresponding to the key if exists.
func (g *Graph) getVertex(k string) (*Vertex, error) {
	for _, v := range g.vertices {
		if v.key == k {
			return v, nil
		}
	}

	return nil, fmt.Errorf(errVertexNotExists, k)
}

// exists checks if a vertex already exists in graph vertices
func (g *Graph) exists(key string) bool {
	for _, v := range g.vertices {
		if v.key == key {
			return true
		}
	}

	return false
}
