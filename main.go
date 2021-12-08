package main

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/Tibart/socialanalogizer/analogizer"
	"github.com/Tibart/socialanalogizer/graph"
)

func main() {

	// Create graph instance
	var g *graph.Graph
	g = new(graph.Graph)

	// Get new vertices, add them to graph and return new vertices
	folder := "./data"
	suffix := ".md"
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatalf("could not read directory: %s", err)
	}
	allVertices := []string{}
	newVertices := []string{}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), suffix) {
			vertex := strings.TrimSuffix(f.Name(), suffix)
			allVertices = append(allVertices, vertex)
			if !g.Containes(vertex) {
				if err := g.AddVertex(vertex); err != nil {
					log.Fatalf("could not add vertex: %s", err)
				} else {
					newVertices = append(newVertices, vertex)
				}
			}
		}
	}

	// Add reference to new vertex in existing vertices documents
	if err := analogizer.Analogize(allVertices, newVertices, g, folder, suffix); err != nil {
		log.Fatalln(err.Error())
	}

	// Print graph
	g.Print()
}
