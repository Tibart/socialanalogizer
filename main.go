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

	// Collect all vertices and correcpomding markdown documents
	folder := "./data"
	suffix := ".md"
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatalf("could not read directory: %s", err)
	}
	allVertices := []string{}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), suffix) {
			vertex := strings.TrimSuffix(f.Name(), suffix)
			allVertices = append(allVertices, vertex)
		}
	}

	// Create graph instance
	var g *graph.Graph = new(graph.Graph)

	// Create Analogizer instance
	a, _ := analogizer.NewAnalogizer(g, allVertices)

	// Update vertex adjacents and amend corresponding markdown document
	for _, vertex := range allVertices {
		// Open file`
		doc, err := os.OpenFile(path.Join(folder, vertex+suffix), os.O_RDWR, 0644)
		if err != nil {
			panic(err)
		}
		defer doc.Close()

		// Amend new verteces
		b, err := a.Amend(vertex, doc)
		if err != nil {
			panic(err)
		}

		// Write amended document
		if _, err := doc.WriteAt(b, 0); err != nil {
			panic(err)
		}
	}

	// Print graph
	a.Print()

}
