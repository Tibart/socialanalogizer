package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"strings"
)

func main() {

	// Add data item's as vertices to the graph structure

	// Scan data objects for occurrences of known vertices.

	// Create hyperlink of occurrences to data objects

	// Export graph to json object for visualization

	// TODO: Implement scanner https://pkg.go.dev/text/scanner

	//tstDoc := []byte("This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. ")

	folder := "./data"
	files, _ := ioutil.ReadDir(folder)

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".md") {
			// Read doc
			d, _ := ioutil.ReadFile(path.Join(folder, f.Name()))

			// scan doc
			fmt.Println(f.Name())
			fmt.Println("----------------------------------")
			fmt.Println("BEFORE:")
			fmt.Print(string(d))
			fmt.Println()
			//tstKeys := []string{"test d"}

			s := "developer"
			expr := fmt.Sprintf(`(?i)\b%s\b`, s)
			re, _ := regexp.Compile(expr)

			fmt.Println("\nAFTER:")
			ns := fmt.Sprintf("**%s**", s)
			fmt.Println(string(re.ReplaceAll(d, []byte(ns))))
			fmt.Println()

		}
	}

	// if re.MatchString(tstDoc) {
	// 	fmt.Printf("Test document value: %v\n", re.ReplaceAllString(tstDoc, " TESTING "))
	// }

	//fmt.Printf("Test document value: %v\n", tstDoc)
}
