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

	// TODO: Implement scanner https://pkg.go.dev/text/scanner

	//tstDoc := []byte("This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. This is a test document that needs to be scanned for keyword's or phrases. ")
	tstDoc, _ := ioutil.ReadFile("./data/developer.md")
	fmt.Println("BEFORE:")
	fmt.Print(string(tstDoc))
	fmt.Println()
	//tstKeys := []string{"test d"}

	s := "developer"
	expr := fmt.Sprintf(`(?i)\b%s\b`, s)
	re, _ := regexp.Compile(expr)

	fmt.Println("AFTER:")
	ns := fmt.Sprintf("**%s**", s)
	fmt.Println(string(re.ReplaceAll(tstDoc, []byte(ns))))

	// if re.MatchString(tstDoc) {
	// 	fmt.Printf("Test document value: %v\n", re.ReplaceAllString(tstDoc, " TESTING "))
	// }

	//fmt.Printf("Test document value: %v\n", tstDoc)
}
