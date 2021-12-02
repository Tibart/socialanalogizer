package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/Tibart/socialanalogizer/graph"
)

func main() {

	// Create graph instance
	g := graph.Graph{}

	// Get new vertices, add them to graph and return new vertices
	folder := "./data"
	suffix := ".md"
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatalf("could not read directory: %s", err)
	}

	nv := []string{}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), suffix) {
			v := strings.TrimSuffix(f.Name(), suffix)
			if !g.Containes(v) {
				if err := g.AddVertex(v); err != nil {
					log.Fatalf("could not add vertex: %s", err)
				} else {
					nv = append(nv, v)
				}
			}
		}
	}

	// Add reference to new vertex in existing vertices documents
	for _, f := range files {
		tv := strings.TrimSuffix(f.Name(), suffix)
		for _, fv := range nv {
			// Exclude new vertex files
			if tv == fv {
				continue
			}

			// TODO: replace by open file
			// Read existing vertex document
			vd, err := ioutil.ReadFile(path.Join(folder, f.Name()))
			if err != nil {
				log.Fatalf("could not read file: %s", err)
			}

			// Find occurrences of vertex key
			expr := fmt.Sprintf(`(?i)(?P<key>\b%s\b)`, fv)
			re, err := regexp.Compile(expr)
			if err != nil {
				log.Fatalf("could not compile regexp: %s", err)
			}

			if re.Match(vd) {
				// Add edge to graph when match is present
				if err := g.AddEdge(fv, tv); err != nil {
					log.Fatalf("could not add edge: %s", err)
				}

				// Replace occurrences with template
				tmpl := []byte("*****> $key <*****")
				result := re.ReplaceAll(vd, tmpl)

				// TODO: use open file instead of reopening file every time
				// Write result back to file
				if err := os.WriteFile(path.Join(folder, f.Name()), result, 0644); err != nil {
					log.Fatalf("could not write to file: %s", err)
				}
			}
		}
	}

	// Print graph
	fmt.Println(g.String())
}
