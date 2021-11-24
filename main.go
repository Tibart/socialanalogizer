package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {

	// Add data item's as vertices to the graph structure

	// Scan data objects for occurrences of known vertices.

	// Create hyperlink of occurrences to data objects

	// Export graph to json object for visualization

	//tstDoc := []byte("This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. ")
	tstDoc, _ := ioutil.ReadFile("./data/developer.md")
	fmt.Print(string(tstDoc))
	fmt.Println()
	//tstKeys := []string{"test d"}

	re, _ := regexp.Compile(`(?i)\btest\b`)

	fmt.Println(string(re.ReplaceAll(tstDoc, []byte("Annemarie"))))

	// if re.MatchString(tstDoc) {
	// 	fmt.Printf("Test document value: %v\n", re.ReplaceAllString(tstDoc, " TESTING "))
	// }

	//fmt.Printf("Test document value: %v\n", tstDoc)
}
