package analogizer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"

	"github.com/Tibart/socialanalogizer/graph"
)

// Init, scan all documents for references and update them if nessecary
func Analogize(allVertices, newVertices []string, graph *graph.Graph, folder, suffix string) error {
	// Add reference to new vertex in existing vertices documents
	for _, toVertex := range allVertices {
		for _, fromVertex := range newVertices {
			// Exclude new vertex files
			if toVertex == fromVertex {
				continue
			}

			// TODO: replace by open file or storage abstraction
			// Read existing vertex document
			doc, err := ioutil.ReadFile(path.Join(folder, toVertex) + suffix)
			if err != nil {
				return fmt.Errorf("could not read file: %s", err)
			}

			// Find occurrences of vertex key
			expr := fmt.Sprintf(`(?i)(?P<key>(?:\[)?\b%s\b(?:\.md\))?)`, fromVertex)
			re, err := regexp.Compile(expr)
			if err != nil {
				return fmt.Errorf("could not compile regexp: %s", err)
			}

			if re.Match(doc) {
				// Add edge to graph when match is present
				if err := graph.AddEdge(fromVertex, toVertex); err != nil {
					return fmt.Errorf("could not add edge: %s", err)
				}

				// Replace occurrences in document
				result := re.ReplaceAllFunc(doc, func(s []byte) []byte {
					// Exclude already markdown hyperlinked references
					if string(s[0]) == "[" || string(s[len(s)-1]) == `)` {
						return s
					}
					// Replace occurrences keyword with markdown hyperlink
					return []byte(fmt.Sprintf(`[%[1]s](./%[1]s.md)`, s))
				})

				// TODO: use open file instead of reopening file every time
				// Write result back to file
				if err := os.WriteFile(path.Join(folder, toVertex)+suffix, result, 0644); err != nil {
					return fmt.Errorf("could not write to file: %s", err)
				}
			}
		}
	}

	return nil
}
