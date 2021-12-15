package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/Tibart/socialanalogizer/analogizer"
	"github.com/Tibart/socialanalogizer/graph"
)

func main() {

	// Create graph instance
	var g *graph.Graph = new(graph.Graph)

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

	// Iterare over files
	for _, vertex := range allVertices {
		// Open file`
		doc, err := os.OpenFile(path.Join(folder, vertex+suffix), os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}
		defer doc.Close()

		// Amend new verteces
		b, err := analogizer.Amend(newVertices, vertex, doc, g)
		if err != nil {
			panic(err)
		}

		// TODO: only write when changed
		// Write amended document
		if _, err := doc.WriteAt(b, 0); err != nil {
			panic(err)
		}
	}

	// Print graph
	g.Print()
}
