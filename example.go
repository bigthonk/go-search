package main

import (
	"fmt"
	"github.com/bigthonk/go-search" 
)

func main() {
	engine := searchengine.New()

	doc1 := searchengine.Document{ID: 0, Title: "Hello World"}
	doc2 := searchengine.Document{ID: 1, Title: "Go is fun"}
	doc3 := searchengine.Document{ID: 2, Title: "Hello Golang"}

	engine.AddDoc(doc1)
	engine.AddDoc(doc2)
	engine.AddDoc(doc3)

	// Normal search
	results := engine.Search("Hello")
	for _, doc := range results {
		fmt.Println(doc.Title)
	}

	// Fuzzy search
	resultsFuzzy := engine.FuzzySearch("Helol", 2)
	for _, doc := range resultsFuzzy {
		fmt.Println(doc.Title)
	}
}