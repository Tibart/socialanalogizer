package analogizer

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/Tibart/socialanalogizer/graph"
)

type Analogizer struct {
	grh *graph.Graph
}

func NewAnalogizer(graph *graph.Graph, vertices []string) (Analogizer, error) {
	a := &Analogizer{}
	a.grh = graph

	// Restate all vertices
	if _, err := a.restate(vertices); err != nil {
		return Analogizer{}, err
	}

	return *a, nil
}

// Restate adds vertices to the graph and returns the nummer of new vertices
func (a *Analogizer) restate(vertices []string) (int, error) {
	c := 0
	for _, v := range vertices {
		if !a.grh.Containes(v) {
			if err := a.grh.AddVertex(v); err != nil {
				return c, err
			} else {
				c++
			}
		}
	}
	return c, nil
}

// Amend new vertices references to document
func (a *Analogizer) Amend(key string, r io.Reader) ([]byte, error) {
	// Read document
	doc, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	orgLen := len(doc)

	// Add vertex to graph
	if !a.grh.Containes(key) {
		if err := a.grh.AddVertex(key); err != nil {
			return nil, err
		}
	}

	// Iterate over all know vertices
	for _, k := range a.grh.GetVerticesKeys() {
		// Exclude new vertex document
		if k == key {
			continue
		}

		// Find occurrences of vertex key in document
		expr := fmt.Sprintf(`(?i)(?P<key>(?:\[)?\b%s\b(?:\.md\))?)`, k)
		re, err := regexp.Compile(expr)
		if err != nil {
			return nil, fmt.Errorf("could not compile regexp: %s", err)
		}

		if re.Match(doc) {
			// Add edge to graph when match is present
			if err := a.grh.AddEdge(k, key, true); err != nil {
				return nil, fmt.Errorf("could not add edge: %s", err)
			}

			// Replace occurrences in document
			doc = re.ReplaceAllFunc(doc, func(s []byte) []byte {
				// Exclude already markdown hyperlinked references
				if string(s[0]) == "[" || string(s[len(s)-1]) == `)` {
					return s
				}

				// Replace occurrences keyword with markdown hyperlink
				hl := html.EscapeString(strings.ReplaceAll(string(s), " ", "%20"))
				return []byte(fmt.Sprintf(`[%[1]s](./%[2]s.md)`, s, hl))
			})
		}
	}

	// Return nil when document did not change
	if orgLen == len(doc) {
		return nil, nil
	}

	return doc, nil
}

func (a *Analogizer) Print() {
	j, _ := json.MarshalIndent(a.grh.Export(), "", "   ")
	fmt.Println(string(j))
}
