package analogizer

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"

	"github.com/Tibart/socialanalogizer/graph"
)

type Analogizer struct {
	grh *graph.Graph
}

func NewAnalogizer(g *graph.Graph) (Analogizer, error) {
	a := &Analogizer{}
	a.grh = g

	return *a, nil
}

// Restate adds vertices to the graph and returns the nummer of new vertices
func (a *Analogizer) Restate(vertices []string) (int, error) {
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
				return []byte(fmt.Sprintf(`[%[1]s](./%[1]s.md)`, s))
			})
		}
	}

	// Return nil when document did not change
	if orgLen == len(doc) {
		return nil, nil
	}

	return doc, nil
}
