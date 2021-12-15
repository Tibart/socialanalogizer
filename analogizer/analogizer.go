package analogizer

import (
	"fmt"
	"io"
	"io/ioutil"
	"regexp"

	"github.com/Tibart/socialanalogizer/graph"
)

// Amend new vertices references to document
func Amend(newVertices []string, toVertex string, doc io.Reader, g *graph.Graph) ([]byte, error) {
	// Read document
	b, err := ioutil.ReadAll(doc)
	if err != nil {
		return nil, err
	}

	for _, v := range newVertices {
		// Exclude new vertex document
		if v == toVertex {
			continue
		}

		// Find occurrences of vertex key in document
		expr := fmt.Sprintf(`(?i)(?P<key>(?:\[)?\b%s\b(?:\.md\))?)`, v)
		re, err := regexp.Compile(expr)
		if err != nil {
			return nil, fmt.Errorf("could not compile regexp: %s", err)
		}

		if re.Match(b) {
			// Add edge to graph when match is present
			if err := g.AddEdge(v, toVertex); err != nil {
				return nil, fmt.Errorf("could not add edge: %s", err)
			}

			// Replace occurrences in document
			b = re.ReplaceAllFunc(b, func(s []byte) []byte {
				// Exclude already markdown hyperlinked references
				if string(s[0]) == "[" || string(s[len(s)-1]) == `)` {
					return s
				}

				// Replace occurrences keyword with markdown hyperlink
				return []byte(fmt.Sprintf(`[%[1]s](./%[1]s.md)`, s))
			})
		}
	}

	return b, nil
}
