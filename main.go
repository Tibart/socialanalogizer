package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/Tibart/socialanalogizer/graph"
)

func main() {

	// Add data item's as vertices to the graph structure

	// Scan data objects for occurrences of known vertices.

	// Create hyperlink of occurrences to data objects

	// Export graph to json object for visualization

	// TODO: Implement scanner https://pkg.go.dev/text/scanner

	//tstDoc := []byte("This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. ")

	// Create graph instance
	g := graph.Graph{}
	g.AddVertex("dolor")

	// Get to be vertices
	folder := "./data"
	suffix := ".md"
	files, _ := ioutil.ReadDir(folder)

	// Create map of new vertices (nv)
	nv := []string{}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), suffix) {
			v := strings.TrimSuffix(f.Name(), suffix)
			if !g.Containes(v) {
				nv = append(nv, v)
			}
		}
	}

	// Add reference to new vertex in existing vertices documents
	for _, f := range files {
		for _, v := range nv {
			// Exclude new vertex
			if strings.TrimSuffix(f.Name(), suffix) != v {
				//doc, _ := ioutil.ReadFile(path.Join(folder, f.Name()))
				// Open vertex file
				file, _ := os.Open(path.Join(folder, f.Name()))
				stats, _ := file.Stat()
				defer file.Close()

				// Read vertex doc
				vd := make([]byte, stats.Size())
				l, _ := file.Read(vd)

				file.Write()

				// expr := fmt.Sprintf(`(?i)\b%s\b`, v)
				// re, _ := regexp.Compile(expr)
				// repl := []byte("*****> $0 <*****")
				// result := re.ReplaceAll(doc, repl)

				// fmt.Println(string(result))

			}

		}
	}

	// for _, f := range files {
	// 	if strings.HasSuffix(f.Name(), ".md") {
	// 		// Read doc
	//

	// 		// scan doc
	// 		fmt.Println(f.Name())
	// 		fmt.Println("----------------------------------")
	// 		fmt.Println("BEFORE:")
	// 		fmt.Print(string(d))
	// 		fmt.Println()
	// 		//tstKeys := []string{"test d"}

	// 		s := "Ipsum"
	// 		expr := fmt.Sprintf(`(?i)\b%s\b`, s)
	// 		re, _ := regexp.Compile(expr)

	// 		fmt.Println("\nAFTER:")
	// 		ns := fmt.Sprintf("**%s**", s)
	// 		fmt.Println(string(re.ReplaceAll(d, []byte(ns))))
	// 		fmt.Println()

	// 	}
	// }

	// if re.MatchString(tstDoc) {
	// 	fmt.Printf("Test document value: %v\n", re.ReplaceAllString(tstDoc, " TESTING "))
	// }

	//fmt.Printf("Test document value: %v\n", tstDoc)
}
